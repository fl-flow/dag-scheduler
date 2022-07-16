package jobscheduler

import (
  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)



func RunningJobs() {
  defer Wait.Done()
  ActionLoop(db.DataBase.Where("status = ?", model.JobRunning).Preload("Tasks"), runningJobOne)
}


func runningJobOne(j model.Job) {
  isSuccess := true
  for _, t := range j.Tasks {
    if t.Status == model.TaskFailed {
      toChangeJobStatus(j, model.JobFailed)
      return
    } else if t.Status == model.TaskTimeout {
      toChangeJobStatus(j, model.JobFailed)
      return
    } else if t.Status == model.TaskCancelled {
      toChangeJobStatus(j, model.JobCancelled)
      return
    } else if t.Status != model.TaskSuccess {
      isSuccess = false
    }
  }
  if isSuccess {
    toChangeJobStatus(j, model.JobSuccess)
  }
}


func toChangeJobStatus(j model.Job, status model.JobStatus) {
  ret := db.DataBase.Debug().Model(&model.Job{ID: j.ID}).Where(
    "status = ?", model.JobRunning,
  ).Updates(model.Job{
    Status: status,
  })
  NotifyJob(
    ret,
    j,
    status,
  )
}
