package form

import (
  "github.com/fl-flow/dag-scheduler/common/parser"
)


type JobCreateForm struct {
  parser.Conf
  Name          string                `json:"name" binding:"required"`
  NotifyUrl     string                `json:"notify_url"`
  JobNotifyUrl  string                `json:"job_notify_url"`
  ID            uint                  `json:"id"`
  WaitCmdToRun  bool                  `json:"wait_cmd_to_run" binding:"required"`
}
