package scheduler

import (
  "log"
  "sync"
  "time"
  "gorm.io/gorm"
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
  log.Println("some loop func error")
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
