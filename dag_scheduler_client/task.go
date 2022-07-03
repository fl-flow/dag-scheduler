package dagschedulerclient

import (
  "log"
  "encoding/json"

  "github.com/fl-flow/dag-scheduler/common/error"
)


func ToRunTask(data interface{}) *error.Error {
  b, err := json.Marshal(data)
  if err != nil {
    log.Fatalf("data json dumps error:  %v\n", err)
  }
  _, e := fetch("POST", "/api/v1/task/torun/", b)
  return e
}
