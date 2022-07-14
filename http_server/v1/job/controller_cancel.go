package job

import (
  "fmt"
  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/error"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/common/operation"
)


func JobCancelController(f JobCancelForm) *error.Error {
  jobID := f.JobID
  ret := db.DataBase.Find(&model.Job{})
  if ret.RowsAffected == 0 {
    return &error.Error{Code: 110020, Hits: fmt.Sprintf("%v", jobID)}
  }
  operation.Cancel(jobID, "cancelled by user")
  return nil
}
