package multiprocessing

import (
  "io"
  "log"
  "fmt"
  "sync"
  "bytes"
  "net/http"
  "mime/multipart"
  "github.com/fl-flow/dag-scheduler/common/resource"
)


type DataStream struct {
  Done      bool
  Data      []byte
}


type Process struct {
  Cmd       string
  Node      string
  Uid       string
  Memory    uint
  Done      *sync.WaitGroup
}


func NewProcess(
  cmd string,
  memory uint,
) (*Process, bool) {
  wait := &sync.WaitGroup{}
  wait.Add(1)
  argWait := &sync.WaitGroup{}
  argWait.Add(1)
  var node, uid string
  go resource.Resource.Allocating(
    memory,
    &node,
    &uid,
    argWait,
    wait,
  )
  argWait.Wait()
  return &Process{
    Cmd: cmd,
    Node: node,
    Uid: uid,
    Memory: memory,
    Done: wait,
  }, true
}


func (p *Process)Run(ch chan DataStream, chOutput chan DataStream) {
  pr, rw := io.Pipe()
  go func() {
    for true {
      d := <-ch
      if d.Done {
        break
      }
      rw.Write(d.Data)
    }
    rw.Close()
  }()
  body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
  writer.WriteField("cmd", p.Cmd)
  w, _ := writer.CreateFormField("data")
  io.Copy(w, pr)
  request, err := http.NewRequest(
    "POST",
    fmt.Sprintf("%v/api/v1/task-process/", p.Node),
    body,
  )
  request.Header.Set("Content-Type", writer.FormDataContentType())
  if err != nil {
      log.Fatal(err)
  }
  response, err := http.DefaultClient.Do(request)
  if err != nil {
    log.Fatal(err)
  }
  for true {
    buf := make([]byte, 1024)
    length, e := response.Body.Read(buf)
    if e == io.EOF {
      if length > 0 {
        chOutput <- DataStream {
          Done: false,
          Data: buf[:length],
        }
      }
      chOutput <- DataStream{
        Done: true,
      }
      p.Done.Done()
      // if !resource.Resource.ResourceNodeDown(p.Node, p.Uid) {
      //   log.Fatal(
      //     fmt.Sprintf("error resource down node: %v, uid: %v", p.Node, p.Uid),
      //   )
      // }
      break
    }
    // TODO: assert http error
    chOutput <- DataStream {
      Done: false,
      Data: buf[:length],
    }
  }
}
