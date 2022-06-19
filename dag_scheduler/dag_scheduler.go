package dagscheduler

import (
  "github.com/fl-flow/dag-scheduler/etc"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/scheduler"
)

func Run()  {
  if etc.IsRunSchedulerLoop {
    scheduler.Loop()
  }
}
