package resource

import (
  "log"

  "github.com/fl-flow/dag-scheduler/etc"
  "github.com/fl-flow/resource-coordinator/client"
)

var ResourceManager *client.ResourceNodeType


func init()  {
  resourceCoordinatorClient := client.Client{
    Schema: "http",
    IP: etc.ResourceCoordinatorIP,
    Port: etc.ResourceCoordinatorPort,
  }
  resource, ret := resourceCoordinatorClient.NewResource("dagscheduler memery");
  if !ret {
    log.Fatalf("error new resource")
  }
  var r bool
  ResourceManager, r = resource.NewResourceNode(
    etc.NodeId,
    etc.SchedulerLoopMemory,
    0,
    0,
  )
  if !r {
    log.Fatalf("error new resource node")
  }
}
