package scheduler

import (
  "log"
  "gorm.io/gorm/clause"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


func InitTasks() {
  defer wait.Done()
  ActionLoop(initTaskOne, "init")
}


func initTaskOne(t model.Task) bool {
  if t.MemoryLimited < (TotalMemory - LockedMemory) {
    MemoryRwMutex.Lock()
    defer MemoryRwMutex.Unlock()
    if t.MemoryLimited > (TotalMemory - LockedMemory) {
      return true
    }
    log.Println("task init -> ready: id-", t.ID)
    LockedMemory = LockedMemory + t.MemoryLimited
    // TODO: database lock (test in mysql)
    tx := db.DataBase.Debug().Begin()
    defer tx.Commit()
    tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&t)
    if t.Status != "init" {
      return false
    }
    t.Status = "ready"
    tx.Save(&t)
    return false
  } else {
    return true
  }
}
