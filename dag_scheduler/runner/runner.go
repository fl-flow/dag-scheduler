package runner

import (
  "io"
  "fmt"
  "sync"
  "bufio"
  "os/exec"
  "encoding/json"
  "encoding/base64"
)


func Run(
  cmd []string,
  commonParameters string,
  parameters interface{},
  inputs []string,
) ([]string, string, bool) {
  var process *exec.Cmd
  if len(cmd) == 1 {
    process = exec.Command(cmd[0])
  } else {
    process = exec.Command(cmd[0], cmd[1:]...)
  }
  stdin, _ := process.StdinPipe()
  stdout, _ := process.StdoutPipe()
  stderror, _ := process.StderrPipe()
  wait := &sync.WaitGroup{}
  wait.Add(1)
  var rets []string
  go inputArgs(stdin, commonParameters, parameters, inputs)
  go getRet(stdout, &rets, wait)
  process.Start()
  errorBytes, _ := io.ReadAll(stderror)
  errorString := string(errorBytes)
  if errorString != "" {
    return rets, fmt.Sprintf(`{"code": error, "error_msg": %v}`, errorString), false
  }
  e := process.Wait()
  if e != nil {
    return rets, fmt.Sprintf(`{"code": %v, "error_msg": %v}`, e, errorString), false
  }
  wait.Wait()
  return rets, "success", true
}


func inputArgs(
  w io.Writer,
  commonParameters string,
  parameters interface {},
  inputs []string,
) {
  write2Pipe(w, commonParameters)
  parametersBytes, _ := json.Marshal(parameters)
  write2Pipe(w, string(parametersBytes))
  for _, i := range inputs {
    write2Pipe(w, i)
  }
}


func write2Pipe(w io.Writer, inputData string) {
  encodedData := base64.StdEncoding.EncodeToString([]byte(inputData))
  w.Write([]byte(encodedData + "\n"))
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
    encodedData, _ := base64.StdEncoding.DecodeString(string(part))
    *rets = append(*rets, string(encodedData))
  }
}
