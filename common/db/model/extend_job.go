package model

import (
	"database/sql/driver"
)


type JobStatus int

const (
	JobInit				JobStatus = 1
	JobRunning		JobStatus = 2
	JobSuccess		JobStatus = 3
	JobFailed			JobStatus = 4
	JobCancelled	JobStatus = 5
)

func (c *JobStatus) Scan(value interface{}) error {
	*c = JobStatus(value.(int64))
  return nil
}

func (c JobStatus) Value() (driver.Value, error) {
  return int64(c), nil
}
