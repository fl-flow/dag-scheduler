package controller

import (
  "fmt"
  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/error"
  "github.com/fl-flow/dag-scheduler/common/parser"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/common/parser/dag_parser"
)


func JobCreate(name string, conf parser.Conf) (model.Job, *error.Error) {
  // parse conf
  var job model.Job
  orderedTasksMap, parameterMap, e := parser.Parse(conf)
  if e != nil {
    return job, e
  }
  // Are there any tasks
  if len(orderedTasksMap) == 0 {
    return job, &error.Error{Code: 110010}
  }

  // TODO: cmd validate

  // insert to db (job and tasks)
  // insert job
  job = model.Job {
    Name: name,
    Dag: orderedTasksMap,
    RawDag: conf.Dag,
    Parameter: model.JobParameter(parameterMap),
    Status: "init",
  }
  tx := db.DataBase.Begin()
  db.DataBase.Create(&job)
  var tasks []model.Task
  var ups []([]string)
  for group, orderedTasks := range orderedTasksMap {
    tasks, ups = mergeTasks(tasks, ups, orderedTasks, group)
  }
  for index, task := range tasks {
    tasks[index].Job = job
    tasks[index].OrderInJob = index
    tasks[index].MemoryLimited = parameterMap[task.Group].Tasks[task.Name].Setting.Resource.Memory
  }
  // db.DataBase.Create(&tasks)
  tasksInsert(tasks, ups)
  tx.Commit()
  job.Tasks = tasks
  return job, nil
}


func mergeTasks(
  fTasks []model.Task,
  fUps []([]string),
  lTasks *[]dagparser.TaskParsered,
  group string,
) ([]model.Task, []([]string)) {
  anchor := 0
  for _, lt := range *lTasks {
    has := false
    var ups []string
    for _, up := range lt.Depandent.Up {
      ups = append(ups, up.UpTask)
    }
    m := model.Task {
      Name: lt.Name,
      Description: "", // TODO:
      // Pid: nil,
      Group: group,
      Status: "init",
      // UpTasks: ups,
      Cmd: lt.Cmd,
      ValidateCmd: lt.ValidateCmd,
    }
    for index, ft := range fTasks {
      if has {
        if lt.Name == ft.Name {
          if anchor < index{
            anchor = index
          }
        } else {
          break
        }
      }else {
        if lt.Name == ft.Name {
          has = true
          fTasks = append(fTasks, model.Task{})
          copy(fTasks[index+1:], fTasks[index:])
          fTasks[index] = m

          fUps = append(fUps, []string{})
          copy(fUps[index+1:], fUps[index:])
          fUps[index] = ups

          if anchor < index{
            anchor = index
          }
        }
      }
    }
    if !has {
      fTasks = append(fTasks, model.Task{})
      copy(fTasks[anchor+1:], fTasks[anchor:])
      fTasks[anchor] = m

      fUps = append(fUps, []string{})
      copy(fUps[anchor+1:], fUps[anchor:])
      fUps[anchor] = ups

      anchor ++
    }
  }
  return fTasks, fUps
}


func tasksInsert(
  tasks []model.Task,
  ups []([]string),
) {
  // i := 0
  fmt.Println(tasks, ups, "upsupsups")
  insertedTaskName2IDMap := map[string]uint{}
  var insertingTaskBatch []model.Task
  var insertingTaskIndex []int
  for index, t := range tasks {
    // check task'ups is inserted
    canInsert := true
    for _, up := range ups[index] {
      canInsert = false
      for insertedTask, _ := range insertedTaskName2IDMap {
          if up == insertedTask {
            canInsert = true
            break
          }
      }
      if !canInsert {
        break
      }
    }

    // insert task and taskLinks
    if canInsert {
      insertingTaskBatch = append(insertingTaskBatch, t)
      insertingTaskIndex = append(insertingTaskIndex, index)
    } else {
      fmt.Println(ups)
      taskInsert(insertingTaskBatch, insertingTaskIndex, insertedTaskName2IDMap, ups)
      insertingTaskBatch = []model.Task{t}
      insertingTaskIndex = []int{index}
    }
  }
  if len(insertingTaskBatch) != 0 {
    taskInsert(insertingTaskBatch, insertingTaskIndex, insertedTaskName2IDMap, ups)
  }
}


func taskInsert(
  insertingTaskBatch []model.Task,
  insertingTaskIndex []int,
  insertedTaskName2IDMap (map[string]uint),
  ups []([]string),
) {
  fmt.Println(insertingTaskBatch, "tasks")
  fmt.Println(insertedTaskName2IDMap, "map")
  fmt.Println(insertingTaskIndex, "index")
  fmt.Println(ups, "uuuu")
  db.DataBase.Debug().Create(insertingTaskBatch)
  insertingLinks := []model.TaskLink{}
  for index, it := range insertingTaskBatch {
    insertedTaskName2IDMap[it.Name] = it.ID
    fmt.Println(ups[insertingTaskIndex[index]], "ups[insertingTaskIndex[index]]ups[insertingTaskIndex[index]]")
    for _, up := range ups[insertingTaskIndex[index]] {
      insertingLinks = append(insertingLinks, model.TaskLink{
        UpId: insertedTaskName2IDMap[up],
        DownId: it.ID,
      })
    }
  }
  if len(insertingLinks) != 0 {
    db.DataBase.Debug().Create(&insertingLinks)
  }
}
