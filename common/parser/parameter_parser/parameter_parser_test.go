package parameterparser

import (
  "fmt"
  "testing"
  "encoding/json"
  "github.com/fl-flow/dag-scheduler/common/error"
)


func testParse(rawParameter string) (Parameter, *error.Error){
  var parameter Parameter
  ok := json.Unmarshal([]byte(rawParameter), &parameter)
  if ok != nil {
    return parameter, &error.Error{Code: 12000}
  }
  parameter1, ok1 := parameter.Parse()
  if ok1 != nil {
    return parameter, ok1
  }
  return parameter1, nil
}


func TestParse(t *testing.T) {
  _, parameterparserok1 := testParse(`
    {
      "common": "CCCC",
      "tasks": ""
    }
  `)
  fmt.Println(parameterparserok1)

  parameterparser, parameterparserok := testParse(`
    {
      "common": "CCCC",
      "tasks": {"a": "z", "f": "d"}
    }
  `)
  fmt.Println(parameterparserok)
  fmt.Println(parameterparser)


  // // parser unit test
  // tasks, parameter, e := parser.Parse(`
  //   {
  //     "dag": {
  //       "a": {
  //         "input": [],
  //         "output": ["d"],
  //         "cmd": "cmd"
  //       }
  //     },
  //     "parameter": {
  //       "common": "CCCC",
  //       "tasks": {"a": "z"}
  //     }
  //   }
  // `)
  // fmt.Println(e)
  // fmt.Println(tasks)
  // fmt.Println(parameter)
}
