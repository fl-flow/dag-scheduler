package scheduler

import (
  "sync"
  // "log"

  // "github.com/fl-flow/dag-scheduler/etc"
  // "github.com/fl-flow/resource-coordinator/client"
)


var wait *sync.WaitGroup
var Limit int = 5000
// var TotalMemory uint64
// var LockedMemory uint64
// var MemoryRwMutex *sync.RWMutex

// var Resource *client.ResourceNodeType


// func init()  {
//   // TotalMemory = etc.SchedulerLoopMemory
//   // LockedMemory = 0
//   // MemoryRwMutex = new(sync.RWMutex)
//
//
//   resourceCoordinatorClient := client.Client{
//     Schema: "http",
//     IP: etc.ResourceCoordinatorIP,
//     Port: etc.ResourceCoordinatorPort,
//   }
//   resource, ret := resourceCoordinatorClient.NewResource("dagscheduler memery");
//   if !ret {
//     log.Fatalf("error new resource")
//   }
//   var r bool
//   Resource, r = resource.NewResourceNode(
//     etc.NodeId,
//     etc.SchedulerLoopMemory,
//     0,
//     0,
//   )
//   if !r {
//     log.Fatalf("error new resource node")
//   }
// }
