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

  // resource coordinator
  resourceCoordinatorIP := flag.String("resourcecoordinatorip", ResourceCoordinatorIP, "resource coordinator ip")
	resourceCoordinatorPort := flag.Int("resourcecoordinatorport", ResourceCoordinatorPort, "resource coordinator port")

  flag.Parse()


  // scheduler server
  SchedulerIp = *schedulerIp
  SchedulerPort = *schedulerPort
  IsRunHttpApi = *isRunHttpApi


  // scheduler loop
  IsRunSchedulerLoop = *isRunSchedulerLoop
  SchedulerLoopMemory = uint64(*schedulerLoopMemoryMB) * 1024 * 1024
  fixSchedulerLoopMemory()

  // resource coordinator
  ResourceCoordinatorIP  = *resourceCoordinatorIP
  ResourceCoordinatorPort  = *resourceCoordinatorPort


  log.Println(fmt.Sprintf(
      `
      // scheduler server
      IsRunHttpApi: %v
      SchedulerIp: %v
      SchedulerPort: %v

      // scheduler loop
      isRunSchedulerLoop: %v
      SchedulerLoopMemory: %v
      ResourceCoordinatorIP: %v
      ResourceCoordinatorPort: %v
      `,
      // scheduler server
      IsRunHttpApi,
      SchedulerIp,
      SchedulerPort,

      // scheduler loop
      IsRunSchedulerLoop,
      SchedulerLoopMemory,

      // resource coordinator
      ResourceCoordinatorIP,
      ResourceCoordinatorPort,
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
