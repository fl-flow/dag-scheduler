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
  if !t.GotCmdToRun {
    return false
  }
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
    ret := qs.Updates(model.Task{
      Status: model.TaskFailed,
      CmdRet: string(b),
    })
    NotifyTask(
      ret,
      t,
      model.TaskFailed,
    )
    return
  }

  toRunTx := db.DataBase.Begin()

  ret := toRunTx.Model(&model.Task{ID: t.ID}).Where(
    "status = ?", model.TaskReady,
  ).Updates(model.Task{
    Status: model.TaskRunning,
  })
  // status is changed
  if ret.RowsAffected == 0 {
    toRunTx.Rollback()
    return
  }
  if t.OrderInJob == 0 {
    toRunTx.Model(&model.Job{ID: t.JobID}).Updates(
      model.Job{Status: model.JobRunning},
    )
  }
  toRunTx.Commit()
  NotifyTask(
    ret,
    t,
    model.TaskRunning,
  )

  rets, description, ok := runner.Run(
    t.ID,
    t.JobID,
    t.Name,
    t.Group,
    []string(t.Dag.Cmd),
    t.CommonParameter,
    t.Parameters,
    inputs,
    len(t.Dag.Output),
  )

  // unlock pre-distributed memory
  MemoryRwMutex.Lock()
  LockedMemory = LockedMemory - t.Parameters.Setting.Resource.Memory
  MemoryRwMutex.Unlock()

  // update task status -> success or failed
  if !ok {
    ret := db.DataBase.Debug().Model(&model.Task{ID: t.ID}).Where(
      "status = ?", model.TaskRunning,
    ).Updates(model.Task{
      Status: model.TaskFailed,
      CmdRet: description,
    })
    NotifyTask(
      ret,
      t,
      model.TaskFailed,
    )
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
    ret := qs_.Updates(model.Task{
      Status: model.TaskFailed,
      CmdRet: string(bt),
    })
    NotifyTask(
      ret,
      t,
      model.TaskFailed,
    )
    return
  }
  retSuccess := qs_.Updates(model.Task{
    Status: model.TaskSuccess,
    CmdRet: description,
  })
  NotifyTask(
    retSuccess,
    t,
    model.TaskSuccess,
  )
}
