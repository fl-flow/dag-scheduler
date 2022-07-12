package operation

import (
  "log"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/federation"
)


func CancelTask(task model.Task, description string) {
  var ts []model.Task
  db.DataBase.Debug().Model(&model.Task{}).Where(
    "status IN ?", []model.TaskStatus{
      model.TaskInit,
      model.TaskReady,
      model.TaskRunning,
    },
  ).Where(
    "job_id = ?", task.JobID,
  ).Find(&ts)
  var initTaskIDs []uint
  for _, t := range ts {
    switch t.Status {
      case model.TaskInit:
        initTaskIDs = append(initTaskIDs, t.ID)
      case model.TaskReady:
        go (federation.Node { // TODO: description
          ID: t.RunOnNode,
        }).Cancel (
          t.JobID,
          t.Group,
          t.Name,
        )
      case model.TaskRunning:
        go (federation.Node { // TODO: description
          ID: t.RunOnNode,
        }).Cancel (
          t.JobID,
          t.Group,
          t.Name,
        )
      default:
        log.Fatalf("error")
    }
  }
  totalInitCount := len(initTaskIDs)
  if totalInitCount != 0 {
    ret := db.DataBase.Debug().Model(&model.Task{}).Where(
      "id IN ?", initTaskIDs,
    ).Where(
      "job_id = ?", task.JobID,
    ).Updates(model.Task{
      Status: model.TaskCancelled,
      CmdRet: description,
    })
    if ret.RowsAffected != int64(totalInitCount) {
      CancelTask(task, description)
    }
  }


  //
  // db.DataBase.Debug().Model(&model.Task{}).Where(
  //   "status = ?", model.TaskInit,
  // ).Where(
  //   "job_id = ?", task.JobID,
  // ).Updates(model.Task{
  //   Status: model.TaskCancelled,
  //   CmdRet: description,
  // })
  //
  // // TODO: Recycling Resource
  // db.DataBase.Debug().Model(&model.Task{}).Where(
  //   "status = ?", model.TaskReady,
  // ).Where(
  //   "job_id = ?", task.JobID,
  // ).Updates(model.Task{
  //   Status: model.TaskCancelled,
  //   CmdRet: description,
  // })
  //
  // // TODO: kill
  // db.DataBase.Debug().Model(&model.Task{}).Where(
  //   "status = ?", model.RunningTasks,
  // ).Where(
  //   "job_id = ?", task.JobID,
  // ).Updates(model.Task{
  //   Status: model.TaskCancelled,
  //   CmdRet: description,
  // })
}
