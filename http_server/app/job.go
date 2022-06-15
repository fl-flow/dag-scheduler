package app

import (
  "fmt"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/parser"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/http_server/form"
  "github.com/fl-flow/dag-scheduler/http_server/controller"
  "github.com/fl-flow/dag-scheduler/http_server/http/mixin"
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
  var jobs []model.Job
  var total int64
  queryset, page, size := mixin.List(context, db.DataBase.Model(model.Job{}).Preload("Tasks"))
  queryset.Find(&jobs)
  queryset.Count(&total)
  response.R(
    context,
    0,
    "success",
    map[string]interface{}{
      "count": total,
      "result": map[string]interface {} {
        "list": jobs,
        "page": page,
        "size": size,
      },
    },
  )
  return
}
