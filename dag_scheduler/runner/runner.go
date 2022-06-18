package runner

import (
  "io"
  "fmt"
  "sync"
  "bufio"
  "os/exec"
  "encoding/base64"
)

// TODO: it is a test
func Run(cmd string, args... string) ([]string, string, bool) {
  fmt.Println(cmd, args)
  process := exec.Command(cmd)
  stdin, _ := process.StdinPipe()
  stdout, _ := process.StdoutPipe()
  // TODO: stderror, _ := process.StderrPipe()
  wait := &sync.WaitGroup{}
  wait.Add(1)
  var rets []string
  go inputArgs(stdin, args)
  go getRet(stdout, &rets, wait)
  process.Start()
  e := process.Wait()
  if e != nil {
    return rets, fmt.Sprintf("%v", e), false
  }
  wait.Wait()
  fmt.Println(rets)
  return rets, "success", true
}


func inputArgs(w io.Writer, args []string) {
  for _, i := range args {
    encodedData := base64.StdEncoding.EncodeToString([]byte(i))
    w.Write([]byte(encodedData + "\n"))
  }
}


func getRet(r io.Reader, rets *[]string, w *sync.WaitGroup) {
  defer w.Done()
  reader := bufio.NewReader(r)
  for true {
    part, prefix, err := reader.ReadLine()
    if err != nil {
      break
    }
    if prefix {
      continue
    }
    *rets = append(*rets, string(part))
  }
}
