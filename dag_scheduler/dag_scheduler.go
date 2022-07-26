package dagscheduler

import (
  "github.com/fl-flow/dag-scheduler/etc"
  // "github.com/fl-flow/dag-scheduler/dag_scheduler/scheduler"
  // "github.com/fl-flow/dag-scheduler/dag_scheduler/scheduler/job"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/scheduler/processing"
)

func Run()  {
  if etc.IsRunSchedulerLoop {
    // go scheduler.Loop()
    // go processingscheduler.Loop()
    // jobscheduler.Loop()
    processingscheduler.Loop()
  }
}
