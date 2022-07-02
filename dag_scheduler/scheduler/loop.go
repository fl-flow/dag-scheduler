package scheduler

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
  wait = &sync.WaitGroup{}
  wait.Add(1)

  // handle init tasks
  go InitTasks()

  // handle ready tasks
  go ReadyTasks()

  // handle running tasks
  go RunningTasks()

  // handle success tasks
  go SuccessTasks()

  // handle failed tasks
  go FailedTasks()

  // handle timeout tasks
  go TimeoutTasks()

  // handle cancelled tasks
  go CancelledTasks()

  log.Println("loop started")
  wait.Wait()
  log.Fatalf("some loop func error")
}


type DoOneModel func(t model.Task) bool


func ActionLoop(qs *gorm.DB, doOneModel DoOneModel) {
  offset := 0
  for true {
    var tasks []model.Task
    qs.Debug().Offset(offset).Limit(Limit).Find(&tasks,)
    for _, t := range tasks {
      isDone := doOneModel(t)
      if isDone {
        offset = 0
        time.Sleep(time.Second *2)
        continue
      }
    }
    offset = offset + Limit
    if len(tasks) == 0 {
      offset = 0
      time.Sleep(time.Second *2)
    }
  }
}


func NotifyTask(ret *gorm.DB, t model.Task, status model.TaskStatus) {
  if ret.RowsAffected == 0 {
    return
  }
  if ret.RowsAffected != 1 {
    log.Fatalf("error update with lock, bug")
    return
  }
  if t.NotifyUrl == "" {
    return
  }
  notify.NotifyStatus(
    t.NotifyUrl,
    int(status),
    "task",
    t.ID,
    map[string]interface{}{
      "job_id": t.JobID,
      "group": t.Group,
      "task": t.Name,
    },
  )
}
