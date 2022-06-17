package dagscheduler

import (
  "sync"
  "github.com/fl-flow/dag-scheduler/dag_scheduler/scheduler"
)

func Run(wait *sync.WaitGroup)  {
  scheduler.Loop(wait)
}
