package scheduler

import (
  "fmt"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


func ReadyTasks() {
  defer wait.Done()
  ActionLoop(readyTaskOne, "ready")
}


func readyTaskOne(t model.Task) bool {
  // TODO: rwlock
  fmt.Println("ready", t)
  return false
}
