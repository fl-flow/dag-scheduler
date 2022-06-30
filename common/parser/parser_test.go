package parser

import (
  "fmt"
  "testing"
  "encoding/json"
  "github.com/fl-flow/dag-scheduler/common/error"
)


func testParse(rawConf string) (Role2TaskParseredList, Role2Parameter, *error.Error){
  var conf Conf
  var tasks Role2TaskParseredList
  var parameters Role2Parameter
  ok := json.Unmarshal([]byte(rawConf), &conf)
  if ok != nil {
    return tasks, parameters, &error.Error{Code: 10000}
  }
  return conf.Parse()
}


func TestParse(t *testing.T) {
  tasks, parameter, e := testParse(`
    {
      "dag": {
        "host": {
          "a": {
            "input": [],
            "output": ["d"],
            "cmd": "cmd"
          }
        }
      },
      "parameter": {
        "host": {
          "common": "CCCC",
          "tasks": {"a": "z"}
        }
      }
    }
  `)
  fmt.Println(e)
  fmt.Println(tasks)
  fmt.Println(parameter)
}
