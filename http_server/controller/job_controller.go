package controller

import (
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
  for group, orderedTasks := range orderedTasksMap {
    tasks = mergeTasks(tasks, orderedTasks, group)
  }
  for index, _ := range tasks {
    tasks[index].Job = job
    tasks[index].OrderInJob = index
  }
  db.DataBase.Create(&tasks)
  tx.Commit()
  job.Tasks = tasks
  return job, nil
}


func mergeTasks(fTasks []model.Task, lTasks *[]dagparser.TaskParsered, group string) []model.Task {
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
      UpTasks: ups,
      MemoryLimited: 0,
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
      anchor ++
    }
  }
  return fTasks
}
