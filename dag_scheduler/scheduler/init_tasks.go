package scheduler

import (
  "log"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


func InitTasks() {
  defer wait.Done()
  ActionLoop(db.DataBase.Where("status = ?", model.TaskInit), initTaskOne)
}


func initTaskOne(t model.Task) bool {
  if t.MemoryLimited < (TotalMemory - LockedMemory) {
    MemoryRwMutex.Lock()
    defer MemoryRwMutex.Unlock()
    ret := db.DataBase.Debug().Model(&model.Task{ID: t.ID}).Where("status = ?", model.TaskInit).Updates(model.Task{Status: model.TaskReady})
    if ret.RowsAffected == 0 {
      return false
    }
    log.Println("task init -> ready: id-", t.ID)
    LockedMemory = LockedMemory + t.MemoryLimited
    return false
  } else {
    return true
  }
}
