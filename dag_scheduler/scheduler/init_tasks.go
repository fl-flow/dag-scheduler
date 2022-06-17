package scheduler

import (
  "log"
  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


func InitTasks() {
  defer wait.Done()
  ActionLoop(initTaskOne, "init")
}


func initTaskOne(t model.Task) bool {
  // TODO: rwlock
  if t.MemoryLimited < (TotalMemory - LockedMemory) {
    // database lock
    log.Println("task init -> ready: id-", t.ID)
    LockedMemory = LockedMemory + t.MemoryLimited
    t.Status = "ready"
    db.DataBase.Save(&t)
    return false
  } else {
    return true
  }
}
