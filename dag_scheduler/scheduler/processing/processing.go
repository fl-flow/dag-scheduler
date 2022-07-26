package processingscheduler

import (
  "os"
  "fmt"
  "log"
  // "bytes"
  "bufio"
  "strings"
  "strconv"
  "encoding/base64"

  "github.com/fl-flow/dag-scheduler/common/multiprocessing"
)


func RunProcessing(){
  defer Wait.Done()
  file, _ := os.OpenFile(PipeFile, os.O_RDWR, os.ModeNamedPipe)
  filew, _ := os.OpenFile(PipeFileW, os.O_RDWR, os.ModeNamedPipe)
  reader := bufio.NewReader(file)
  // TODO: semaphore
  for true {
    b := readLine(reader)
    rets := strings.SplitN(b, ",", 2)
    cmdEncodedData, _ := base64.StdEncoding.DecodeString(rets[0])
    cmd := string(cmdEncodedData)
    memoryEncodedData, _ := base64.StdEncoding.DecodeString(rets[1])
    memory_ := string(memoryEncodedData)
    memory, err := strconv.ParseInt(memory_, 10, 64)
    if err != nil {
      filew.WriteString("fail\n")
      continue
    }
    p, success := multiprocessing.NewProcess(
      cmd,
      uint(memory),
    )
    if !success {
      filew.WriteString("fail\n")
      continue
    }
    filew.WriteString("success\n")
    ch := make(chan multiprocessing.DataStream)
    chOutput := make(chan multiprocessing.DataStream)
    go inputData(reader, ch)
    go p.Run(ch, chOutput)
    outputData(filew, chOutput)
  }
}


func inputData(reader *bufio.Reader, ch chan multiprocessing.DataStream) {
  for true {
    d := readLine(reader)
    if d == "" {
      fmt.Println("end")
      break
    }
    size, e := strconv.ParseInt(d, 10, 64)
    if e != nil {
      log.Fatal("some error size d ", d)
      break
    }
    buf := make([]byte, size)
    length, _ := reader.Read(buf)
    if int64(len(buf)) != size {
      log.Fatal("some error size ", length, int64(len(buf)))
    }
    ch <- multiprocessing.DataStream{
        Done: false,
        Data: buf,
    }
  }
  ch <- multiprocessing.DataStream{
      Done: true,
  }
}


func outputData(
  filew *os.File,
  outputStream chan multiprocessing.DataStream,
) {
  w := bufio.NewWriter(filew)
  for true {
    d := <- outputStream
    if d.Done {
      break
    }
    b := string(d.Data)
    lengthString := strconv.FormatInt(int64(len(b)), 10)
    w.WriteString(lengthString + "\n")
    w.WriteString(b)
  }
  w.WriteString("\n")
  w.Flush()
}
