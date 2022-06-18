package model

import (
	"database/sql/driver"
)


type TaskStatus	int

const (
	TaskInit			TaskStatus = 1
	TaskReady			TaskStatus = 2
	TaskRunning		TaskStatus = 3
	TaskSuccess		TaskStatus = 4
	TaskFailed		TaskStatus = 5
	TaskTimeout		TaskStatus = 6
	TaskCancelled	TaskStatus = 7
)

func (c *TaskStatus) Scan(value interface{}) error {
	*c = TaskStatus(value.(int64))
  return nil
}

func (c TaskStatus) Value() (driver.Value, error) {
  return int64(c), nil
}
