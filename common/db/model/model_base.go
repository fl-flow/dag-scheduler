package model

import (
	"encoding/json"
	"database/sql/driver"
	"github.com/fl-flow/dag-scheduler/common/parser"
	"github.com/fl-flow/dag-scheduler/common/parser/dag_parser"
	"github.com/fl-flow/dag-scheduler/common/parser/parameter_parser"
)


func getReverseMap(m map[int]string) map[string]int {
	var rm = map[string]int {}
	for v, d := range m {
		rm[d] = v
	}
	return rm
}


type JobDag 				map[string](*([]dagparser.TaskParsered))

func (c JobDag) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JobDag) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}


type JobRawDagmap   map[string](parser.GroupConf)

func (c JobRawDagmap) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JobRawDagmap) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}


type JobParameter   map[string]parameterparser.Parameter

func (c JobParameter) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JobParameter) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}


type TaskUpTasks 		[]string

func (c TaskUpTasks) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *TaskUpTasks) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}
