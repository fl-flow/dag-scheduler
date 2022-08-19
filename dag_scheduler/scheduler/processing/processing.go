package processingscheduler

import (
  "net"
  "fmt"
  "log"
  "bufio"
  "strings"
  "strconv"
  "encoding/base64"

  "github.com/fl-flow/dag-scheduler/common/multiprocessing"

  "github.com/fl-flow/dag-scheduler/etc"
)


func RunProcessing(con net.Conn){
  reader := bufio.NewReader(con)
  writer := bufio.NewWriter(con)
  b := readLine(reader)
  rets := strings.SplitN(b, ",", 2)
  cmdEncodedData, _ := base64.StdEncoding.DecodeString(rets[0])
  cmd := string(cmdEncodedData)
  memoryEncodedData, _ := base64.StdEncoding.DecodeString(rets[1])
  memory_ := string(memoryEncodedData)
  memory, err := strconv.ParseInt(memory_, 10, 64)
  if err != nil {
    writer.WriteString("fail\n")
    writer.Flush()
    return
  }
  p, success := multiprocessing.NewProcess(
    cmd,
    uint(memory),
  )
  if !success {
    writer.WriteString("fail\n")
    writer.Flush()
    return
  }
  writer.WriteString("success\n")
  writer.Flush()
  ch := make(chan multiprocessing.DataStream)
  chOutput := make(chan multiprocessing.DataStream)
  go inputData(reader, ch)
  go p.Run(ch, chOutput)
  outputData(con, chOutput)
  con.Close()
}



func RunServer() {
  defer Wait.Done()
  listener, err :=net.Listen(
    "tcp",
    fmt.Sprintf("%v:%v", etc.MultiprocessIp, etc.MultiprocessPort),
  )
  if err != nil {
    log.Fatal("error process server start: %v", err)
  }
  for {
    conn,e := listener.Accept()
    if e != nil {
      log.Fatal("error process accept %v", e)
    }
    go RunProcessing(conn)
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
  con net.Conn,
  outputStream chan multiprocessing.DataStream,
) {
  w := bufio.NewWriter(con)
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
