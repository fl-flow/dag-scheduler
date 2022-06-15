package main

import (
  _ "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/http_server"
)


func main() {
  httpserver.Run()
}
