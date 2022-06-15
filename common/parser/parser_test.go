package parser

import (
  "fmt"
  "testing"
  "encoding/json"
  "github.com/fl-flow/dag-scheduler/common/error"
  "github.com/fl-flow/dag-scheduler/common/parser/dag_parser"
  "github.com/fl-flow/dag-scheduler/common/parser/parameter_parser"
)


func testParse(rawConf string) ([]dagparser.TaskParsered, parameterparser.Parameter, *error.Error){
  var conf Conf
  var tasks []dagparser.TaskParsered
  var parameters parameterparser.Parameter
  ok := json.Unmarshal([]byte(rawConf), &conf)
  if ok != nil {
    return tasks, parameters, &error.Error{Code: 10000}
  }
  return Parse(conf)
}


func TestParse(t *testing.T) {
  tasks, parameter, e := testParse(`
    {
      "dag": {
        "a": {
          "input": [],
          "output": ["d"],
          "cmd": "cmd"
        }
      },
      "parameter": {
        "common": "CCCC",
        "tasks": {"a": "z"}
      }
    }
  `)
  fmt.Println(e)
  fmt.Println(tasks)
  fmt.Println(parameter)
}
