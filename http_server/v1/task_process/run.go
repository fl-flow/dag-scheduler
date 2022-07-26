package taskprocess

import (
  "io"
  "fmt"
  "sync"
  "bytes"
  "bufio"
  "os/exec"
  "strings"
)


func RunProcess(
  cmd string,
  ch chan ProcessRetChanDataType,
  chRes chan ProcessRetChanDataType,
) {
  // TODO: move to other project path
  var process *exec.Cmd
  cmdArray := strings.Split(cmd, " ")
  if len(cmdArray) == 1 {
    process = exec.Command(cmdArray[0])
  } else {
    process = exec.Command(cmdArray[0], cmdArray[1:]...)
  }
  stdin, _ := process.StdinPipe()
  stdout, _ := process.StdoutPipe()
  stderror, _ := process.StderrPipe()
  wait := &sync.WaitGroup{}
  wait.Add(1)
  go func () {
    for true {
      fmt.Println("s")
      d := <-ch
      fmt.Println("e")
      fmt.Println(d.Done, "dddd")
      if d.Done {
        stdin.Close()
        return
      }
      fmt.Println(string(d.Data), "d.Datad.Datad.Data")
      stdin.Write(d.Data)
    }
  }()
  go func () {
    defer wait.Done()
    // TODO: change to read chunked
    reader := bufio.NewReader(stdout)
    for true {
      part, prefix, err := reader.ReadLine()
      if err != nil {
        break
      }
      if prefix {
        continue
      }
      chRes <- ProcessRetChanDataType{
        Data: bytes.Join([][]byte{part, []byte("\n")}, []byte("")),
        Done: false,
      }
    }
  }()
  process.Start()
  errorBytes, _ := io.ReadAll(stderror)
  e := process.Wait()
  if e != nil {
    errorString := string(errorBytes)
    if errorString != "" {
      chRes <- ProcessRetChanDataType{
        Data: []byte(fmt.Sprintf(`{"code": error, "error_msg": %v}`, errorString)),
        Done: true,
      }
      return
    }
    chRes <- ProcessRetChanDataType{
      Data: []byte(fmt.Sprintf(`{"code": %v, "error_msg": %v}`, e, e)),
      Done: true,
    }
    return
  }
  wait.Wait()
  chRes <- ProcessRetChanDataType{
    Done: true,
  }
}
