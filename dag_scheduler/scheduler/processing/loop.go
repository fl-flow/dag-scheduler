package processingscheduler

import (
  "log"
  "sync"
)

var Wait *sync.WaitGroup


func Loop() {
  log.Println("process loop starting ...")
  Wait = &sync.WaitGroup{}
  Wait.Add(1)
  go RunServer()

  log.Println("loop started")
  Wait.Wait()
  log.Fatalf("some process loop func error")
}
