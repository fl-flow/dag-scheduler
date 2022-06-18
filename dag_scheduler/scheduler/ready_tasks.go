package scheduler

import (
  "fmt"
  "log"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/runner"
)


func ReadyTasks() {
  defer wait.Done()
  ActionLoop(
    db.DataBase.Where(
      "tasks.status = ? AND (up_tasks.status = ? OR up_tasks.status is NULL )",
      model.TaskReady,
      model.TaskSuccess,
    ).Joins(
      "left join task_links on tasks.id = task_links.down_id",
    ).Joins("left join tasks as up_tasks on up_tasks.id = task_links.up_id"),
    readyTaskOne,
  )
}


func readyTaskOne(t model.Task) bool {
  // TODO: run controller switch
  log.Println("got ready task: id-", t.ID)
  go RunReadyTask(t)
  return false
}


func RunReadyTask(t model.Task) {
  ret := db.DataBase.Debug().Model(&model.Task{ID: t.ID}).Where(
    "status = ?", model.TaskReady,
  ).Updates(model.Task{
    Status: model.TaskRunning,
  })
  // status is changed
  if ret.RowsAffected == 0 {
    return
  }
  // TODO: get args
  rets, description, ok := runner.Run(t.Cmd, "ASDASD", "DDD")
  fmt.Println(rets, description, ok)

  // unlock pre-distributed memory
  MemoryRwMutex.Lock()
  LockedMemory = LockedMemory - t.MemoryLimited
  MemoryRwMutex.Unlock()

  // update task status -> success or failed
  qs := db.DataBase.Debug().Model(&model.Task{ID: t.ID}).Where(
    "status = ?", model.TaskRunning,
  )
  if !ok {
    qs.Updates(model.Task{
      Status: model.TaskFailed,
      CmdRet: description,
    })
    return
  }
  qs.Updates(model.Task{
    Status: model.TaskSuccess,
    CmdRet: description,
  })

  // TODO: save rets
}
