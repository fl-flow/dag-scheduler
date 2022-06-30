package parser

import (
  "github.com/fl-flow/dag-scheduler/common/parser/dag_parser"
  "github.com/fl-flow/dag-scheduler/common/parser/parameter_parser"
)


// raw
type Role2DagTaskMap map[string]dagparser.DagTaskMap

type Role2Parameter map[string]parameterparser.Parameter


type Conf struct {
  Dag         Role2DagTaskMap         `json:"dag"`
  Parameter   Role2Parameter          `json:"parameter"`
}


// parsered
type Role2TaskParseredList map[string](*dagparser.TaskParseredList)
