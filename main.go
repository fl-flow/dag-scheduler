package main

import (
  _ "github.com/fl-flow/dag-scheduler/etc"
  _ "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/http_server"
  "github.com/fl-flow/dag-scheduler/dag_scheduler"
)


func main() {
  go dagscheduler.Run()
  httpserver.Run()
}
