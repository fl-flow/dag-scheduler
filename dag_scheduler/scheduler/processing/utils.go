package processingscheduler

import (
  "fmt"
  "log"
  "bufio"
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