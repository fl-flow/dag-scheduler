package scheduler

import (
  "log"
  "sync"
  "fmt"
  "time"
  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


func Loop(w *sync.WaitGroup) {
  defer w.Done()
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


func ActionLoop(doOneModel DoOneModel, filterStatus model.TaskStatus) {
  offset := 0
  for true {
    fmt.Println(filterStatus)
    var tasks []model.Task
    db.DataBase.Model(model.Task{}).Debug().Offset(offset).Limit(Limit).Find(
      &tasks, model.Task{Status: filterStatus},
    )
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
