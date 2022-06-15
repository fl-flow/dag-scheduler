package mixin

import (
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"
)


func List(context *gin.Context, d *gorm.DB) (*gorm.DB, int, int) {
  var pagination PageNumberPagination
  context.ShouldBindQuery(&pagination)
  size := pagination.Size
  if size == 0 {
    size = DefaultSize
  }else if size > MaxSize {
    size = MaxSize
  }
  d_ := d.Offset((pagination.Page - 1) * size).Limit(size)
  return d_, pagination.Page, size
}
