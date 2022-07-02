package jobscheduler

import (
  "log"
  "sync"
  "time"
  "gorm.io/gorm"

  "github.com/fl-flow/dag-scheduler/util/notify"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


func Loop() {
  log.Println("loop starting ...")
  Wait = &sync.WaitGroup{}
  Wait.Add(1)

  // handle running jobs
  go RunningJobs()

  log.Println("loop started")
  Wait.Wait()
  log.Fatalf("some loop func error")
}


type DoOneModel func(j model.Job)


func ActionLoop(qs *gorm.DB, doOneModel DoOneModel) {
  offset := 0
  for true {
    var jobs []model.Job
    qs.Debug().Offset(offset).Limit(Limit).Find(&jobs,)
    for _, j := range jobs {
      doOneModel(j)
    }
    offset = offset + Limit
    if len(jobs) == 0 {
      offset = 0
      time.Sleep(time.Second *2)
    }
  }
}


func NotifyJob(ret *gorm.DB, j model.Job, status model.JobStatus) {
  if ret.RowsAffected == 0 {
    return
  }
  if ret.RowsAffected != 1 {
    log.Fatalf("error update with lock, bug")
    return
  }
  if j.NotifyUrl == "" {
    return
  }
  notify.NotifyStatus(
    j.NotifyUrl,
    int(status),
    "job",
    j.ID,
    nil,
  )
}
