package parser

import (
  "github.com/fl-flow/dag-scheduler/common/parser/dag_parser"
  "github.com/fl-flow/dag-scheduler/common/parser/parameter_parser"
)

type Conf struct {
  Dag         map[string](dagparser.DagTask)  `json:"dag"`
  Parameter   parameterparser.Parameter       `json:"parameter"`
}
