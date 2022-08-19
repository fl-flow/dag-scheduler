package etc

import (
  "log"
  "fmt"
  "flag"

  "github.com/shirou/gopsutil/mem"
)


func init() {
  // scheduler server
  isRunHttpApi := flag.Bool("httpapi", IsRunHttpApi, "is run http api")
  schedulerIp := flag.String("ip", SchedulerIp, "scheduler server ip")
	schedulerPort := flag.Int("port", SchedulerPort, "scheduler server port")

  // scheduler loop
  isRunSchedulerLoop := flag.Bool("schedulerloop", IsRunSchedulerLoop, "is run scheduler loop")
  schedulerLoopMemoryMB := flag.Int("schedulerloopmemory", 0, "scheduler loop memory")

  // distributed multiprocess
  multiprocessIp := flag.String("multiprocessip", MultiprocessIp, "multiprocess server ip")
  multiprocessPort := flag.Int("multiprocessport", MultiprocessPort, "multiprocess server port")
  
  // resource coordinator
  resourceCoordinatorIP := flag.String("resourcecoordinatorip", ResourceCoordinatorIP, "resource coordinator ip")
	resourceCoordinatorPort := flag.Int("resourcecoordinatorport", ResourceCoordinatorPort, "resource coordinator port")

  // cluster
  nodeId := flag.String("nodeid", NodeId, "node id")

  flag.Parse()


  // scheduler server
  SchedulerIp = *schedulerIp
  SchedulerPort = *schedulerPort
  IsRunHttpApi = *isRunHttpApi


  // scheduler loop
  IsRunSchedulerLoop = *isRunSchedulerLoop
  SchedulerLoopMemory = uint64(*schedulerLoopMemoryMB) * 1024 * 1024
  fixSchedulerLoopMemory()

  // distributed multiprocess
  MultiprocessIp = *multiprocessIp
  MultiprocessPort = *multiprocessPort

  // resource coordinator
  ResourceCoordinatorIP  = *resourceCoordinatorIP
  ResourceCoordinatorPort  = *resourceCoordinatorPort

  NodeId = *nodeId

  log.Println(fmt.Sprintf(
      `
      // scheduler server
      IsRunHttpApi: %v
      SchedulerIp: %v
      SchedulerPort: %v

      // scheduler loop
      isRunSchedulerLoop: %v
      SchedulerLoopMemory: %v

      // distributed multiprocess
      multiprocessIp
      multiprocessPort

      // resource coordinator
      ResourceCoordinatorIP: %v
      ResourceCoordinatorPort: %v

      // cluster
      NodeId: %v
      `,
      // scheduler server
      IsRunHttpApi,
      SchedulerIp,
      SchedulerPort,

      // scheduler loop
      IsRunSchedulerLoop,
      SchedulerLoopMemory,

      // distributed multiprocess
      MultiprocessIp,
      MultiprocessPort,

      // resource coordinator
      ResourceCoordinatorIP,
      ResourceCoordinatorPort,

      // cluster
      NodeId,
  ),)
}


func fixSchedulerLoopMemory() {
  memInfo, err := mem.VirtualMemory()
  if err != nil {
    log.Fatalf("error get sys memory")
  }
  totalMemory, freeMemory := memInfo.Total, memInfo.Available
  if SchedulerLoopMemory == 0 {
    log.Println(fmt.Sprintf(
        `
        system total memory: %v
        system free memory: %v
        `,
        totalMemory,
        freeMemory,
      ))
    SchedulerLoopMemory = freeMemory
    return
  }
  if SchedulerLoopMemory > freeMemory {
    log.Fatalf(fmt.Sprintf(
        `
        SchedulerLoopMemory(%v) is lagger than free memory of system
        `,
        SchedulerLoopMemory,
      ))
  }
}
