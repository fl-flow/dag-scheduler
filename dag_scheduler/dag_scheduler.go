package dagscheduler

import (
  "github.com/fl-flow/dag-scheduler/etc"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/scheduler"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/scheduler/job"
)

func Run()  {
  if etc.IsRunSchedulerLoop {
    go scheduler.Loop()
    jobscheduler.Loop()
  }
}
