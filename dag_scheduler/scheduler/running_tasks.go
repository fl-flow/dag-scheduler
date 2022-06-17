package scheduler

import (
  "fmt"
  "time"
)


func RunningTasks() {
  defer wait.Done()
  for true {
    fmt.Println("RunningTasks")
    time.Sleep(time.Second *2)
  }
}
