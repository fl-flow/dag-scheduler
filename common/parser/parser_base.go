package parser

import (
  "github.com/fl-flow/dag-scheduler/common/parser/dag_parser"
  "github.com/fl-flow/dag-scheduler/common/parser/parameter_parser"
)


type GroupConf map[string](dagparser.DagTask)

type Conf struct {
  Dag         map[string]GroupConf                    `json:"dag"`
  Parameter   map[string]parameterparser.Parameter    `json:"parameter"`
}
