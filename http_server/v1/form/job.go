package form

import (
  "github.com/fl-flow/dag-scheduler/common/parser"
)


type JobCreateForm struct {
  parser.Conf
  Name          string                `json:"name" binding:"required"`
}
