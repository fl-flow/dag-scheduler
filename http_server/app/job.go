package app

import (
  "fmt"
  "github.com/gin-gonic/gin"

  "dag/common/db"
  "dag/common/parser"
  "dag/common/db/model"
  "dag/http_server/form"
  "dag/http_server/controller"
  "dag/http_server/http/mixin"
  "dag/http_server/http/response"
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
  queryset := mixin.List(context, db.DataBase.Model(model.Job{}))
  queryset.Find(&jobs)
  queryset.Count(&total)
  response.R(
    context,
    0,
    "success",
    map[string]interface{}{
      "count": total,
      "result": jobs,
    },
  )
  return
}
