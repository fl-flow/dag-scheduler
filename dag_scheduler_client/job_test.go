package dagschedulerclient

import (
  "fmt"
  "testing"
  "encoding/json"
)


func TestCreat(t *testing.T) {
  var jobConf interface{}
  rawJobConf := `
    {
      "name": "你好",
      "dag": {
          "a": {
              "input": [],
              "output": ["d"],
              "cmd": "cmd",
              "validate_cmd": "validate_cmd"
          },
          "b": {
              "input": ["a.d"],
              "output": [],
              "cmd": "cmd"
          }
      },
      "parameter": {
          "common": "CCCC",
          "tasks": {"a": "z", "b": "asd"}
      }
    }
  `
  json.Unmarshal([]byte(rawJobConf), &jobConf)
  // job, e := CreateJob(jobConf)
  fmt.Println(jobConf)
}
