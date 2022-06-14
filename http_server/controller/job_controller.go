package controller

import (
  "dag/common/db"
  "dag/common/error"
  "dag/common/parser"
  "dag/common/db/model"
)


func JobCreate(name string, conf parser.Conf) (model.Job, *error.Error) {
  // parse conf
  var job model.Job
  orderedTasks, _, e := parser.Parse(conf)
  if e != nil {
    return job, e
  }
  // Are there any tasks
  if len(orderedTasks) == 0 {
    return job, &error.Error{Code: 110010}
  }

  // TODO: cmd validate

  // insert to db (job and tasks)
  // insert job
  job = model.Job {
    Name: name,
    Dag: orderedTasks,
    RawDag: conf.Dag,
    Parameter: model.JobParameter(conf.Parameter),
    Status: "init",
  }
  tx := db.DataBase.Begin()
  db.DataBase.Create(&job)
  // insert tasks
  var tasks []model.Task
  for index, t := range orderedTasks {
    var ups []string
    for _, up := range t.Depandent.Up {
      ups = append(ups, up.UpTask)
    }
    tasks = append(tasks, model.Task {
      Job: job,
      Name: t.Name,
      Description: "", // TODO:
      // Pid: nil,
      Status: "init",
      OrderInJob: index,
      UpTasks: ups,
      MemoryLimited: 0, // TODO:
      Cmd: t.Cmd,
      ValidateCmd: t.ValidateCmd,
    })
  }
  db.DataBase.Create(&tasks)
  tx.Commit()
  job.Tasks = tasks
  return job, nil
}
