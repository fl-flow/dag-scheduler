package dagparser

import(
  "fmt"
  "strings"
  "github.com/fl-flow/dag-scheduler/common/error"
)


func (dagMap DagTaskMap) Parse () (TaskParseredList, *error.Error ){
  var tasksParsed TaskParseredList
  tasksDepandentMap, inDegreeMap, error := dagMap.findTasksDepandent()
  if error != nil {
    return tasksParsed, error
  }
  orderedTasks, loopE := tasksDepandentMap.checkLoop(inDegreeMap)
  if loopE != nil {
    return tasksParsed, loopE
  }
  tasksParsed = buildTaskParsered(
    tasksDepandentMap,
    orderedTasks,
    dagMap,
  )
  return tasksParsed, nil
}


func (dagTaskMap DagTaskMap) findTasksDepandent() (
  TasksDepandentMap,
  map[string]int,
  *error.Error,
) {

  // get all tasks
  tasksDepandentMap := make(TasksDepandentMap)
  inDegreeMap := make(map[string]int)
  for taskName, _ := range dagTaskMap {
    tasksDepandentMap[taskName] = &TaskDepandent{}
  }

  // build depandents
  for taskName, taskInfo := range dagTaskMap {
    if len(taskInfo.Cmd) == 0 {
      return tasksDepandentMap, inDegreeMap, &error.Error{
          Code: 11040,
          Hits: taskName,
      }
    }
    input := taskInfo.Input
    for _, inputItem := range input {
      upTaskName, upTag, annotation, e := parseTaskDepandent(inputItem)
      if e != nil {
        return tasksDepandentMap, inDegreeMap, e
      }
      _, inputOk := tasksDepandentMap[upTaskName]
      if !inputOk {
        return tasksDepandentMap, inDegreeMap, &error.Error{
            Code: 11020,
            Hits: fmt.Sprintf(
                "parser error( %v; task %v not exits )",
                inputItem,
                upTaskName,
            ),
        }
      }
      tasksDepandentMap[taskName].Up = append(
        tasksDepandentMap[taskName].Up,
        TaskInput {
          UpTask: upTaskName,
          Tag: upTag,
          Annotation: annotation,
        })
      tasksDepandentMap[upTaskName].Down = append(tasksDepandentMap[upTaskName].Down, taskName)
    }
  }
  for taskName, taskInfo := range tasksDepandentMap {
    inDegreeMap[taskName] = len(taskInfo.Up)
  }
  return tasksDepandentMap, inDegreeMap, nil
}



func parseTaskDepandent(value string) (string, string, string, *error.Error) {
  // task.tag
  rets := strings.SplitN(value, ".", 2)
  if (len(rets) != 2){
    return "", "", "", &error.Error{
        Code: 11010,
        Hits: value,
    }
  }
  tagAnnotation := strings.SplitN(rets[1], ":", 2)
  if len(tagAnnotation) != 2 {
    return rets[0], rets[1], "", nil
  }
  return rets[0], tagAnnotation[0], tagAnnotation[1], nil
}


func (tasksDepandentMap TasksDepandentMap) checkLoop(inDegreeMap map[string]int) ([]string, *error.Error) {
  var queue, orderedTasks []string
  for taskName, inDegree := range inDegreeMap {
    if (inDegree == 0) {
      queue = append(queue, taskName)
    }
  }
  totals := 0
  qLength := len(queue)
  // TODO: order of task
  for (qLength > 0) {
    taskName := queue[0]
    orderedTasks = append(orderedTasks, taskName)
    queue = queue[1:]
    totals ++
    qLength --
    for _, downTaskName := range tasksDepandentMap[taskName].Down {
      inDegreeMap[downTaskName] --
      if inDegreeMap[downTaskName] == 0 {
        queue = append(queue, downTaskName)
        qLength ++
      }
    }
  }
  if (totals != len(tasksDepandentMap)){
    // TODO: find loop
    return orderedTasks, &error.Error{
        Code: 11030,
    }
  }
  return orderedTasks, nil
}


func buildTaskParsered(
  tasksDepandentMap TasksDepandentMap,
  orderedTasks []string,
  dagTaskMap map[string]DagTask,
) TaskParseredList {
    var tasksParsed TaskParseredList
    for _, taskName := range orderedTasks {
      tasksParsed = append(
        tasksParsed,
        TaskParsered {
          Name: taskName,
          Depandent: *(tasksDepandentMap[taskName]),
          Output: dagTaskMap[taskName].Output,
          Cmd: dagTaskMap[taskName].Cmd,
          ValidateCmd: dagTaskMap[taskName].ValidateCmd,
        },
      )
    }
    return tasksParsed
}
