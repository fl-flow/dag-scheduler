package processingscheduler

import (
  "os"
  "fmt"
  "log"
  "bufio"
  "syscall"
)


func readLine(reader *bufio.Reader) string {
  part, prefix, err := reader.ReadLine()
  if err != nil {
    log.Fatal(fmt.Sprintf("error process read line1 %v", err))
  }
  if prefix {
    log.Fatal(fmt.Sprintf("error process read line2 %v", prefix))
  }
  return string(part)
}


func initPipe(){
  os.Remove(PipeFile)
  os.Remove(PipeFileW)
  err := syscall.Mkfifo(PipeFile, 0666)
  if err != nil {
    log.Fatal("create named pipe error:", err)
  }
  errw := syscall.Mkfifo(PipeFileW, 0666)
  if errw != nil {
    log.Fatal("create named pipe error:", errw)
  }
}
