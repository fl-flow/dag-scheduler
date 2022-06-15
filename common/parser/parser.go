package parser

import (
  "fmt"
  "github.com/fl-flow/dag-scheduler/common/error"
  "github.com/fl-flow/dag-scheduler/common/parser/dag_parser"
  "github.com/fl-flow/dag-scheduler/common/parser/parameter_parser"
)


func Parse(conf Conf) (map[string](*([]dagparser.TaskParsered)), map[string](parameterparser.Parameter), *error.Error){
  // var tasks []dagparser.TaskParsered
  var allTasksMap = make(map[string](*([]dagparser.TaskParsered)))
  var allParametersMap = make(map[string](parameterparser.Parameter))
  for group, groupConf := range conf.Dag {
    tasks, dagerror := dagparser.Parse(groupConf)
    if dagerror != nil {
      return allTasksMap, allParametersMap, dagerror
    }
    allTasksMap[group] = &tasks
  }

  for group, groupParameter := range conf.Parameter{
    parameters, parameterError := parameterparser.Parse(groupParameter)
    if parameterError != nil {
      return allTasksMap, allParametersMap, parameterError
    }
    allParametersMap[group] = parameters
  }
  error := checkDagParameter(allTasksMap, allParametersMap)
  if error != nil {
    return allTasksMap, allParametersMap, error
  }
  return allTasksMap, allParametersMap, nil
}


func checkDagParameter(
  taskMap map[string](*([]dagparser.TaskParsered)),
  parameterMap map[string](parameterparser.Parameter)) *error.Error {
  for group, tasks := range taskMap {
    p, ok := parameterMap[group]
    if !ok {
      return &error.Error{
        Code: 11021,
        Hits: group,
      }
    }
    e := checkOneceDagParameter(*tasks, p)
    if e != nil {
      return e
    }
  }
  return nil
}


func checkOneceDagParameter(
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
