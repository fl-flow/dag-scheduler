package etc

import (
  "log"
  "fmt"
  "flag"
)


func init() {
  // scheduler server
  isRunHttpApi := flag.Bool("httpapi", IsRunHttpApi, "is run http api")
  schedulerIp := flag.String("ip", SchedulerIp, "scheduler server ip")
	schedulerPort := flag.Int("port", SchedulerPort, "scheduler server port")

  // scheduler loop
  isRunSchedulerLoop := flag.Bool("schedulerloop", IsRunSchedulerLoop, "is run scheduler loop")

  flag.Parse()

  // scheduler server
  SchedulerIp = *schedulerIp
  SchedulerPort = *schedulerPort
  IsRunHttpApi = *isRunHttpApi

  // scheduler loop
  IsRunSchedulerLoop = *isRunSchedulerLoop

  log.Println(fmt.Sprintf(
      `
      // scheduler server
      IsRunHttpApi: %v
      SchedulerIp: %v
      SchedulerPort: %v

      // scheduler loop
      isRunSchedulerLoop: %v
      `,
      // scheduler server
      IsRunHttpApi,
      SchedulerIp,
      SchedulerPort,

      // scheduler loop
      IsRunSchedulerLoop,
  ),)
}
