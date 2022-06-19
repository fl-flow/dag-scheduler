package scheduler

import (
  "sync"

  "github.com/fl-flow/dag-scheduler/etc"
)


var wait *sync.WaitGroup
var Limit int = 5000
var TotalMemory uint64
var LockedMemory uint64
var MemoryRwMutex *sync.RWMutex


func init()  {
  TotalMemory = etc.SchedulerLoopMemory
  LockedMemory = 0
  MemoryRwMutex = new(sync.RWMutex)
}
