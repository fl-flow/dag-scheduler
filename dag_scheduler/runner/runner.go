package runner

import (
  "io"
  "log"
  "fmt"
  "sync"
  "bufio"
  "strconv"
  "os/exec"
  "encoding/json"
  "encoding/base64"

  "github.com/fl-flow/dag-scheduler/dag_scheduler/tracker"
)


type RunningHookType func(pid int)


func Run(
  taskID uint,
  jobID uint,
  taskName string,
  group string,
  cmd []string,
  commonParameters string,
  parameters interface{},
  inputs []tracker.Input,
  outputLength int,
  runningHook RunningHookType,
) (string, []string, string, bool) {
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
  var summary string
  go inputArgs(
    stdin,
    taskID,
    jobID,
    taskName,
    group,
    commonParameters,
    parameters,
    inputs,
    outputLength,
  )
  go getRet(stdout, &summary, &rets, wait)
  process.Start()
  runningHook(process.Process.Pid)
  errorBytes, _ := io.ReadAll(stderror)
  errorString := string(errorBytes)
  if errorString != "" {
    return summary, rets, fmt.Sprintf(`{"code": error, "error_msg": %v}`, errorString), false
  }
  e := process.Wait()
  if e != nil {
    return summary, rets, fmt.Sprintf(`{"code": %v, "error_msg": %v}`, e, errorString), false
  }
  wait.Wait()
  return summary, rets, "success", true
}


func inputArgs(
  w io.Writer,
  taskID uint,
  jobID uint,
  taskName string,
  group string,
  commonParameters string,
  parameters interface {},
  inputs []tracker.Input,
  outputLength int,
) {
  taskInfo, _ := json.Marshal(map[string]string{
    "job_id": strconv.Itoa(int(jobID)),
    "task_id": strconv.Itoa(int(taskID)),
    "group": group,
    "task_name": taskName,
  })
  write2Pipe(w, string(taskInfo))
  write2Pipe(w, commonParameters)
  parametersBytes, _ := json.Marshal(parameters)
  write2Pipe(w, string(parametersBytes))
  write2Pipe(w, strconv.Itoa(len(inputs)))
  write2Pipe(w, strconv.Itoa(outputLength))
  for _, i := range inputs {
    input, _ := json.Marshal(i)
    write2Pipe(w, string(input))
  }
}


func write2Pipe(w io.Writer, inputData string) {
  encodedData := base64.StdEncoding.EncodeToString([]byte(inputData))
  w.Write([]byte(encodedData + "\n"))
}


func getRet(r io.Reader, summary *string, rets *[]string, w *sync.WaitGroup) {
  defer w.Done()
  reader := bufio.NewReader(r)
  part, prefix, err := reader.ReadLine()
  if err != nil {
    log.Fatalln("error get summary err")
  }
  if prefix {
    log.Fatalln("error get summary prefix")
  }
  encodedData, _ := base64.StdEncoding.DecodeString(string(part))
  *summary = string(encodedData)
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
