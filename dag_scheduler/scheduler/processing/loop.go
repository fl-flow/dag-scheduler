package processingscheduler

import (
  "log"
  "sync"
)

const PipeFile = "/Users/xinchengshao/Desktop/workspace/dag-scheduler/process_pipe.ipc"
const PipeFileW = "/Users/xinchengshao/Desktop/workspace/dag-scheduler/process_pipe_w.ipc"
var Wait *sync.WaitGroup


func Loop() {
  log.Println("process loop starting ...")
  Wait = &sync.WaitGroup{}
  Wait.Add(1)
  initPipe()
  go RunProcessing()

  log.Println("loop started")
  Wait.Wait()
  log.Fatalf("some process loop func error")
}
