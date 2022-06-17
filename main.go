package main

import (
  "sync"
  _ "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/http_server"
  "github.com/fl-flow/dag-scheduler/dag_scheduler"
)


func main() {
  wait := &sync.WaitGroup{}
  wait.Add(1)
  go dagscheduler.Run(wait)
  go httpserver.Run(wait)
  wait.Wait()
}
