package operation


// func CancelTask(task model.Task, description string) {
//   db.DataBase.Debug().Model(&model.Task{}).Where(
//     "status = ?", model.TaskInit,
//   ).Where(
//     "job_id = ?", task.JobID,
//   ).Updates(model.Task{
//     Status: model.TaskCancelled,
//     CmdRet: description,
//   })
//
//   // TODO: Recycling Resource
//   db.DataBase.Debug().Model(&model.Task{}).Where(
//     "status = ?", model.TaskReady,
//   ).Where(
//     "job_id = ?", task.JobID,
//   ).Updates(model.Task{
//     Status: model.TaskCancelled,
//     CmdRet: description,
//   })
//
//   // TODO: kill
//   db.DataBase.Debug().Model(&model.Task{}).Where(
//     "status = ?", model.RunningTasks,
//   ).Where(
//     "job_id = ?", task.JobID,
//   ).Updates(model.Task{
//     Status: model.TaskCancelled,
//     CmdRet: description,
//   })
// }
