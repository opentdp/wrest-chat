package crond

import (
	"errors"
	"strings"

	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"
	"github.com/robfig/cron/v3"

	"github.com/opentdp/wrest-chat/dbase/cronjob"
	"github.com/opentdp/wrest-chat/dbase/tables"
	"github.com/opentdp/wrest-chat/wclient"
	"github.com/opentdp/wrest-chat/wclient/aichat"
	"github.com/opentdp/wrest-chat/wclient/deliver"
)

var crontab *cron.Cron
var logger *logman.Logger

func Daemon() {

	logger = logman.Named("cronjob")
	logger.Info("cron:daemon start")

	crontab = cron.New(cron.WithSeconds())

	jobs, err := cronjob.FetchAll(&cronjob.FetchAllParam{})
	if err != nil || len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		AttachJob(job)
	}

	crontab.Start()

}

// 触发计划任务

func Execute(id uint) error {

	job, _ := cronjob.Fetch(&cronjob.FetchParam{Rd: id})

	if job.Content == "" {
		return errors.New("content is empty")
	}
	if job.Deliver == "-" {
		return errors.New("deliver is empty")
	}

	logger.Info("cron:run "+job.Name, "entryId", job.EntryId)

	// 发送文本内容
	if job.Type == "TEXT" {
		return deliver.Send(job.Deliver, job.Content)
	}

	// 发送AI生成的文本
	if job.Type == "AI" {
		wc := wclient.Register()
		if wc == nil {
			logger.Error("cron:ai", "error", "wclient is nil")
			return errors.New("wclient is nil")
		}
		self := wc.CmdClient.GetSelfInfo()
		data := aichat.Text(job.Content, self.Wxid, "")
		return deliver.Send(job.Deliver, data)
	}

	// 执行命令获取结果
	output, err := command.Exec(&command.ExecPayload{
		Name:          "cron: " + job.Name,
		CommandType:   job.Type,
		WorkDirectory: job.Directory,
		Content:       job.Content,
		Timeout:       job.Timeout,
	})
	if err != nil {
		logger.Warn("cron:run "+job.Name, "error", err)
		return err
	}

	// 发送命令执行结果
	logger.Warn("cron:run "+job.Name, "output", output)
	if output != "" {
		return deliver.Send(job.Deliver, output)
	}

	return nil

}

// 激活计划任务

func AttachJob(job *tables.Cronjob) error {

	cmd := func(id uint) func() {
		return func() { Execute(id) }
	}

	sepc := []string{
		job.Second, job.Minute, job.Hour, job.DayOfMonth, job.Month, job.DayOfWeek,
	}

	entryId, err := crontab.AddFunc(strings.Join(sepc, " "), cmd(job.Rd))
	if err != nil {
		return err
	}

	logger.Info("cron:attach "+job.Name, "entryId", entryId)
	err = cronjob.Update(&cronjob.UpdateParam{
		Rd:      job.Rd,
		EntryId: int64(entryId),
	})

	return err

}

// 管理生命周期

func NewById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		AttachJob(job)
	}

}

func UndoById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		logger.Info("cron:remove "+job.Name, "entryId", job.EntryId)
		crontab.Remove(cron.EntryID(job.EntryId))
	}

}

func RedoById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		logger.Info("cron:update "+job.Name, "entryId", job.EntryId)
		crontab.Remove(cron.EntryID(job.EntryId))
		AttachJob(job)
	}

}

// 获取执行状态

type JobStatus struct {
	EntryId  int64 `json:"entry_id"`
	NextTime int64 `json:"next_time"`
	PrevTime int64 `json:"prev_time"`
}

func GetEntries() map[uint]JobStatus {

	list := map[uint]JobStatus{}

	jobs, err := cronjob.FetchAll(&cronjob.FetchAllParam{})
	if err != nil || len(jobs) == 0 {
		return list
	}

	for _, job := range jobs {
		entry := crontab.Entry(cron.EntryID(job.EntryId))
		list[job.Rd] = JobStatus{
			EntryId:  int64(entry.ID),
			NextTime: entry.Next.Unix(),
			PrevTime: entry.Prev.Unix(),
		}
	}

	return list

}
