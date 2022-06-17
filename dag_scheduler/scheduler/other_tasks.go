package scheduler

import (
  "fmt"
  "time"
)


func SuccessTasks() {
  defer wait.Done()
  for true {
    fmt.Println("SuccessTasks")
    time.Sleep(time.Second *2)
  }
}


func FailedTasks() {
  defer wait.Done()
  for true {
    fmt.Println("FailedTasks")
    time.Sleep(time.Second *2)
  }
}


func TimeoutTasks() {
  defer wait.Done()
  for true {
    fmt.Println("TimeoutTasks")
    time.Sleep(time.Second *2)
  }
}


func CancelledTasks() {
  defer wait.Done()
  for true {
    fmt.Println("CancelledTasks")
    time.Sleep(time.Second *2)
  }
}
