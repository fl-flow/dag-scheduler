package scheduler

import (
  "log"
  "encoding/json"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/runner"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/tracker"
)


func ReadyTasks() {
  defer wait.Done()
  ActionLoop(
    db.DataBase.Where(
      "tasks.status = ?",
      model.TaskReady,
    ).Preload("UpTasks"),
    readyTaskOne,
  )
}


func readyTaskOne(t model.Task) bool {
  for _, upTask := range t.UpTasks {
    if upTask.Status != model.TaskSuccess {
      return false
    }
  }
  // TODO: run controller switch
  log.Println("got ready task: id-", t.ID)
  go RunReadyTask(t)
  return false
}


func RunReadyTask(t model.Task) {
  qs := db.DataBase.Debug().Model(&model.Task{ID: t.ID}).Where(
    "status = ?", model.TaskReady,
  )

  inputs, error := tracker.GetInput(t)
  if error != nil {
    b, _ := json.Marshal(error.Message())
    qs.Updates(model.Task{
      Status: model.TaskFailed,
      CmdRet: string(b),
    })
    return
  }

  ret := qs.Updates(model.Task{
    Status: model.TaskRunning,
  })
  // status is changed
  if ret.RowsAffected == 0 {
    return
  }

  rets, description, ok := runner.Run(
    []string(t.Dag.Cmd),
    t.CommonParameter,
    t.Parameters,
    inputs,
  )

  // unlock pre-distributed memory
  MemoryRwMutex.Lock()
  LockedMemory = LockedMemory - t.Parameters.Setting.Resource.Memory
  MemoryRwMutex.Unlock()

  // update task status -> success or failed
  if !ok {
    db.DataBase.Debug().Model(&model.Task{ID: t.ID}).Where(
      "status = ?", model.TaskRunning,
    ).Updates(model.Task{
      Status: model.TaskFailed,
      CmdRet: description,
    })
    return
  }

  tx := db.DataBase.Begin()
  defer tx.Commit()
  e := tracker.SaveOutput(t, rets)
  qs_ := tx.Debug().Model(&model.Task{ID: t.ID}).Where(
    "status = ?", model.TaskRunning,
  )
  if e != nil {
    bt, _ := json.Marshal(e.Message())
    qs_.Updates(model.Task{
      Status: model.TaskFailed,
      CmdRet: string(bt),
    })
    return
  }
  qs_.Updates(model.Task{
    Status: model.TaskSuccess,
    CmdRet: description,
  })
}
