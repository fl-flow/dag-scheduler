package main

import (
  "flag"
  
  _ "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/http_server"
  "github.com/fl-flow/dag-scheduler/dag_scheduler"
)


func main() {
  ip := flag.String("ip", "127.0.0.1", "ip")
	port := flag.Int("port", 8000, "port")
	flag.Parse()

  go dagscheduler.Run()
  httpserver.Run(*ip, *port)
}
