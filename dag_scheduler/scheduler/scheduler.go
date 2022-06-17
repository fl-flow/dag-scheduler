package scheduler

import (
  "sync"
)


var wait *sync.WaitGroup
var Limit int = 5000
var TotalMemory int64
var LockedMemory int64


func init()  {
  TotalMemory = 1 // TODO: read from conf; or get from computer info
  LockedMemory = 0
}
