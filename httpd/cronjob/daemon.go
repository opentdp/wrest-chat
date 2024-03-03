package cronjob

import (
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"
	"github.com/robfig/cron/v3"

	"github.com/opentdp/wechat-rest/dbase/cronjob"
	"github.com/opentdp/wechat-rest/dbase/tables"
)

var crontab *cron.Cron
var logger *logman.Logger

func Daemon() {

	crontab = cron.New(cron.WithSeconds())
	logger = logman.Named("cronjob")

	jobs, err := cronjob.FetchAll(&cronjob.FetchAllParam{})
	if err != nil || len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		AttachJob(job)
	}

	crontab.Start()

}

func AttachJob(job *tables.Cronjob) error {

	sepc := job.Second + " " + job.Minute + " " + job.Hour + " " + job.DayOfMonth + " " + job.Month + " " + job.DayOfWeek

	task := func() {
		logger.Info("Cron:run " + job.Name)
		result, err := command.Exec(&command.ExecPayload{
			Name:          "Cron: " + job.Name,
			CommandType:   job.Type,
			WorkDirectory: job.Directory,
			Content:       job.Content,
			Timeout:       job.Timeout,
		})
		if err != nil {
			logger.Error("Cron:run "+job.Name, "error", err)
		} else {
			logger.Info("Cron:run "+job.Name, "result", result)
		}
	}

	logger.Info("Cron:add " + job.Name)
	entryId, err := crontab.AddFunc(sepc, task)
	if err != nil {
		return err
	}

	err = cronjob.Update(&cronjob.UpdateParam{
		Rd:      job.Rd,
		EntryId: int64(entryId),
	})

	return err

}

func NewById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		AttachJob(job)
	}

}

func UndoById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		crontab.Remove(cron.EntryID(job.EntryId))
	}

}

func RedoById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		crontab.Remove(cron.EntryID(job.EntryId))
		AttachJob(job)
	}

}

// 获取任务执行状态

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
