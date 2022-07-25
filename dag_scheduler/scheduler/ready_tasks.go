package scheduler

import (
  "fmt"
  "log"
  "encoding/json"

  "github.com/fl-flow/dag-scheduler/etc"
  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/common/resource"
  "github.com/fl-flow/dag-scheduler/common/operation"
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
  if t.RunOnNode != etc.NodeId {
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
    operation.Cancel(t.JobID, "") // // TODO: description
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

  summary, rets, description, ok := runner.Run(
    t.ID,
    t.JobID,
    t.Name,
    t.Group,
    []string(t.Dag.Cmd),
    t.CommonParameter,
    t.Parameters.Args,
    inputs,
    len(t.Dag.Output),
    func(pid int) {
      db.DataBase.Model(&model.Task{ID: t.ID}).Updates(
        model.Task{
          Pid: pid,
        },
      )
    },
  )

  // unlock pre-distributed memory
  // MemoryRwMutex.Lock()
  // LockedMemory = LockedMemory - t.Parameters.Setting.Resource.Memory
  // MemoryRwMutex.Unlock()
  if !resource.ResourceManager.ResourceNodeDown(fmt.Sprintf("%v", t.ID)) {
    log.Fatalf("error reset resource error 2")
  }

  // update task status -> success or failed
  if !ok {
    ret := db.DataBase.Debug().Model(&model.Task{ID: t.ID}).Where(
      "status = ?", model.TaskRunning,
    ).Updates(model.Task{
      Status: model.TaskFailed,
      CmdRet: description,
    })
    operation.Cancel(t.JobID, "") // // TODO: description
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
    operation.Cancel(t.JobID, "") // // TODO: description
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
    Summary: json.RawMessage(summary),
  })
  NotifyTask(
    retSuccess,
    t,
    model.TaskSuccess,
  )
}
