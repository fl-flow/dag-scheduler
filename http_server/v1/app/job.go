package app

import (
  "fmt"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/parser"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/http_server/v1/form"
  "github.com/fl-flow/dag-scheduler/http_server/http/mixin"
  "github.com/fl-flow/dag-scheduler/http_server/v1/controller"
  "github.com/fl-flow/dag-scheduler/http_server/http/response"
)


func JobCreate(context *gin.Context) {
  f := form.JobCreateForm{}
	if e := context.ShouldBindJSON(&f); e != nil {
    response.R(
      context,
      100,
      fmt.Sprintf("%v", e),
      fmt.Sprintf("%v", e),
    )
    return
	}
  job, error := controller.JobCreate(
    f.Name,
    parser.Conf {
      Dag: f.Dag,
      Parameter: f.Parameter,
    },
  )
  if error != nil {
    response.R(
      context,
      error.Code,
      error.Message(),
      error.Message(),
    )
    return
  }
  response.R(context, 0, "success", job)
}


func JobList(context *gin.Context) {
  qs, total, page, size := mixin.List(
    context,
    db.DataBase.Model(&model.Job{}),
  )
  var jobs []model.Job
  qs.Preload("Tasks").Find(&jobs)
  mixin.ListResponse(
    context,
    jobs,
    total,
    page,
    size,
  )
  return
}
