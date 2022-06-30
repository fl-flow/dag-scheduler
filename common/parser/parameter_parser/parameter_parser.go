package parameterparser

import (
  "github.com/fl-flow/dag-scheduler/common/error"
)


func (parameter Parameter) Parse () (Parameter, *error.Error) {
  return parameter, nil
}
