package scheduler

import (
  "fmt"
  "log"

  "github.com/fl-flow/dag-scheduler/etc"
  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/common/resource"
)


func InitTasks() {
  defer wait.Done()
  ActionLoop(db.DataBase.Where("status = ?", model.TaskInit), initTaskOne)
}


func initTaskOne(t model.Task) bool {
  mem := t.Parameters.Setting.Resource.Memory
  if !resource.ResourceManager.ResourceNodeUp(
    mem,
    fmt.Sprintf("%v", t.ID),
  ) {
    return true
  }

  // if mem < (TotalMemory - LockedMemory) {
  //   MemoryRwMutex.Lock()
  //   defer MemoryRwMutex.Unlock()
  ret := db.DataBase.Debug().Model(
    &model.Task{ID: t.ID},
  ).Where("status = ?", model.TaskInit).Updates(
    model.Task{
      Status: model.TaskReady,
      RunOnNode: etc.NodeId,
    },
  )
  if ret.RowsAffected == 0 {
    if !resource.ResourceManager.ResourceNodeDown(fmt.Sprintf("%v", t.ID)) {
      log.Fatalf("error reset resource error")
    }
    return false
  }
  NotifyTask(
    ret,
    t,
    model.TaskReady,
  )
  log.Println("task init -> ready: id-", t.ID)
  // LockedMemory = LockedMemory + mem
  return false
  // }
}
