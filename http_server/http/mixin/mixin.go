package mixin

import (
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/dag-scheduler/http_server/http/response"
)


func List(context *gin.Context, d *gorm.DB, listContainer interface{}) {
  var pagination PageNumberPagination
  var total int64
  d.Count(&total)
  context.ShouldBindQuery(&pagination)
  size := pagination.Size
  if size <= 0 {
    size = DefaultSize
  }else if size > MaxSize {
    size = MaxSize
  }
  page := pagination.Page
  if page <= 0 {
    page = 1
  }
  d_ := d.Offset((page - 1) * size).Limit(size)
  d_.Find(&listContainer)
  response.R(
    context,
    0,
    "success",
    map[string]interface{}{
      "count": total,
      "list": listContainer,
      "page": page,
      "size": size,
    },
  )
}
