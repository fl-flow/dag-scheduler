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


type CmdType	[]string

func (c CmdType) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *CmdType) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}


type JobDag parser.Role2TaskParseredList

func (c JobDag) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JobDag) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}


type JobRawDagmap   parser.Role2DagTaskMap

func (c JobRawDagmap) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JobRawDagmap) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}


type JobParameter   parser.Role2Parameter

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


type TaskDag 	dagparser.TaskParsered

func (c TaskDag) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *TaskDag) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}


type TaskParameter parameterparser.TaskParameter

func (c TaskParameter) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *TaskParameter) Scan(src any) error {
	return json.Unmarshal(([]byte)(src.(string)), c)
}
