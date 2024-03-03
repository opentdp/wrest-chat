package cronjob

import (
	"github.com/opentdp/go-helper/command"
	"github.com/robfig/cron/v3"

	"github.com/opentdp/wechat-rest/dbase/cronjob"
	"github.com/opentdp/wechat-rest/dbase/tables"
)

var crontab *cron.Cron

func Daemon() {

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

func AttachJob(job *tables.Cronjob) error {

	sepc := job.Second + " " + job.Minute + " " + job.Hour + " " + job.DayofMonth + " " + job.Month + " " + job.DayofWeek
	task := func() {
		command.Exec(&command.ExecPayload{
			Name:          "Cron: " + job.Name,
			CommandType:   job.Type,
			WorkDirectory: job.Directory,
			Content:       job.Content,
			Timeout:       job.Timeout,
		})
	}

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

func GetEntries(jobs []*tables.Cronjob) map[uint]any {

	list := map[uint]any{}

	for _, job := range jobs {
		entry := crontab.Entry(cron.EntryID(job.EntryId))
		list[job.Rd] = map[string]any{
			"EntryId":  entry.ID,
			"NextTime": entry.Next.Unix(),
			"PrevTime": entry.Prev.Unix(),
		}
	}

	return list

}
