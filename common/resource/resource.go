package resource

import (
  "log"

  "github.com/fl-flow/dag-scheduler/etc"
  "github.com/fl-flow/resource-coordinator/client"
)


var Resource *client.ResourceType
var ResourceManager *client.ResourceNodeType


func init()  {
  resourceCoordinatorClient := client.Client{
    Schema: "http",
    IP: etc.ResourceCoordinatorIP,
    Port: etc.ResourceCoordinatorPort,
  }
  var ret bool
  Resource, ret = resourceCoordinatorClient.NewResource("dagscheduler memery");
  if !ret {
    log.Fatalf("error new resource")
  }
  var r bool
  ResourceManager, r = Resource.NewResourceNode(
    etc.NodeId,
    etc.SchedulerLoopMemory,
    0,
    0,
  )
  if !r {
    log.Fatalf("error new resource node")
  }
}
