package task

import (
  "fmt"
  "log"
  "syscall"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/common/resource"
)


func TaskCancelController(f TaskCancelForm) (string, bool) {
  var t model.Task
  ret := db.DataBase.Model(&model.Task{}).Where(
    "job_id=? AND `group`=? AND name=?",
    f.JobID,
    f.Group,
    f.Task,
  ).Find(&t)
  if ret.RowsAffected == 0 {
    return "it is not existed", false
  }
  // TODO: switch and move to controller
  if (t.Status == model.TaskInit) {
    ret := db.DataBase.Model(&model.Task{}).Where(
      "job_id=? AND `group`=? AND name=? AND status = ?",
      f.JobID,
      f.Group,
      f.Task,
      model.TaskInit,
    ).Updates(model.Task{
      Status: model.TaskCancelled,
    })
    if ret.RowsAffected == 1 {
      return "success", true
    } else if ret.RowsAffected != 0 {
      log.Fatalf("error")
    }
  }
  if (t.Status == model.TaskReady) {
    // if t.Pid == 0 {
    //   return "failed", false
    // }
    ret := db.DataBase.Model(&model.Task{}).Where(
      "job_id=? AND `group`=? AND name=? AND status = ?",
      f.JobID,
      f.Group,
      f.Task,
      model.TaskReady,
    ).Updates(model.Task{
      Status: model.TaskCancelled,
    })
    if ret.RowsAffected == 1 {
      resource.ResourceManager.ResourceNodeDown(fmt.Sprintf("%v", t.ID))
      return "success", true
    } else if ret.RowsAffected != 0 {
      log.Fatalf("error")
    }
  }
  if t.Status == model.TaskRunning {
    e := syscall.Kill(t.Pid, syscall.SIGTERM)
    if e != nil {
      return "failed", false
    }
    ret :=db.DataBase.Model(&model.Task{}).Where(
      "job_id=? AND `group`=? AND name=? AND status = ?",
      f.JobID,
      f.Group,
      f.Task,
      model.TaskRunning,
    ).Updates(model.Task{
      Status: model.TaskCancelled,
    })
    if ret.RowsAffected == 1 {
      return "success", true
    } else if ret.RowsAffected != 0 {
      log.Fatalf("error")
    }
  }
  return "failed", false
}
