package job

import (
  "fmt"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/parser"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/http_server/http/mixin"
  "github.com/fl-flow/dag-scheduler/http_server/http/response"
)


func JobCreateView(context *gin.Context) {
  f := JobCreateForm{}
	if e := context.ShouldBindJSON(&f); e != nil {
    response.R(
      context,
      100,
      fmt.Sprintf("%v", e),
      fmt.Sprintf("%v", e),
    )
    return
	}
  job, error := JobCreateController(
    f,
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


func JobListView(context *gin.Context) {
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
