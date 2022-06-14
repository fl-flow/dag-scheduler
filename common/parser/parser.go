package parser

import (
  "fmt"
  "dag/common/error"
  "dag/common/parser/dag_parser"
  "dag/common/parser/parameter_parser"
)


func Parse(conf Conf) ([]dagparser.TaskParsered, parameterparser.Parameter, *error.Error){
  var tasks []dagparser.TaskParsered
  var parameters parameterparser.Parameter

  tasks, dagerror := dagparser.Parse(conf.Dag)
  if dagerror != nil {
    return tasks, parameters, dagerror
  }
  parameters, parameterError := parameterparser.Parse(conf.Parameter)
  if parameterError != nil {
    return tasks, parameters, parameterError
  }
  error := checkDagParameter(tasks, parameters)
  if error != nil {
    return tasks, parameters, error
  }
  return tasks, parameters, nil
}


func checkDagParameter(
  tasks []dagparser.TaskParsered,
  parameters parameterparser.Parameter) *error.Error {
  if len(tasks) != len(parameters.Tasks) {
    return &error.Error{Code: 12010}
  }
  for _, task := range tasks {
    _, ok := parameters.Tasks[task.Name]
    if !ok {
      return &error.Error{
        Code: 12020,
        Hits: fmt.Sprintf(
            "dag's task %v is not in parameters",
            task.Name,
        ),
      }
    }
  }
  return nil
}
