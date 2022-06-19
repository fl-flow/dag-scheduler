package etc

import (
  "log"
  "flag"
)


func init() {
  schedulerIp := flag.String("ip", SchedulerIp, "ip")
	schedulerPort := flag.Int("port", SchedulerPort, "port")

  flag.Parse()

  SchedulerIp = *schedulerIp
  SchedulerPort = *schedulerPort

  log.Println("\nSchedulerIp: ", SchedulerIp, "\nSchedulerPort: ", SchedulerPort)
}
