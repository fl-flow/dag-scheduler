package task

import (
  "fmt"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/http_server/http/response"
)


func TaskRunning(context *gin.Context) {
  f := TaskRunningForm{}
	if e := context.ShouldBindJSON(&f); e != nil {
    response.R(
      context,
      100,
      fmt.Sprintf("%v", e),
      fmt.Sprintf("%v", e),
    )
    return
	}
  ret := db.DataBase.Model(&model.Task{}).Where(
    "job_id=? AND `group`=? AND name=?",
    f.JobID,
    f.Group,
    f.Task,
  ).Updates(model.Task{
    GotCmdToRun: true,
  })
  if ret.RowsAffected == 0 {
    response.R(
      context,
      100,
      "it is not existed",
      "it is not existed",
    )
    return
  }
  response.R(context, 0, "success", map[string]string{})
}


func TaskCancelView(context *gin.Context) {
  f := TaskCancelForm{}
	if e := context.ShouldBindJSON(&f); e != nil {
    response.R(
      context,
      100,
      fmt.Sprintf("%v", e),
      fmt.Sprintf("%v", e),
    )
    return
	}
  msg, ret := TaskCancelController(f)
  if !ret {
    response.R(context, 100, msg, map[string]string{})
    return
  }
  response.R(context, 0, "success", map[string]string{})
}
