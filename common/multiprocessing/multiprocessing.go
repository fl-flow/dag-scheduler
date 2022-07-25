package multiprocessing

import (
  "fmt"
  "github.com/fl-flow/dag-scheduler/common/resource"
)


type Process struct {
  Cmd       string
  Node      string
  Uid       string
  Memory    uint
}


func NewProcess(
  cmd string,
  memory uint,
) (*Process, bool) {
  node, uid, success := resource.Resource.Alloc(memory)
  if !success {
    return nil, success
  }
  return &Process{
    Cmd: cmd,
    Node: node,
    Uid: uid,
    Memory: memory,
  }, true
}


func (p *Process)Run(ch chan string) {
  for true {
    done := <-ch
    if done == "true" {
      break
    }
    d := <-ch
    fmt.Println(d, "ddddd")
  }
  fmt.Println("zzzzz", p.Node)
}
