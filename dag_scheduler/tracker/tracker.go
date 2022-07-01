package tracker

import(
  "fmt"
  "strings"

  "github.com/fl-flow/dag-scheduler/common/db"
  "github.com/fl-flow/dag-scheduler/common/error"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


type Input struct {
  Value       string    `json:"value"`
  Annotation  string    `json:"annotation"`
}


func GetInput(t model.Task) ([]Input, *error.Error) {
  var input []Input
  var upTaskIds []uint
  for _, upTask := range t.UpTasks {
    upTaskIds = append(upTaskIds, upTask.ID)
  }
  if len(upTaskIds) != 0 {
    var tagQuerys []string
    for _, up := range t.Dag.Depandent.Up {
      tagQuerys = append(tagQuerys, fmt.Sprintf(`tag = "%v"`, up.Tag))
    }
    tagQuery := strings.Join(tagQuerys, " OR ")
    var upTaskOutputs []model.TaskResult
    db.DataBase.Debug().Where("task_id IN ?", upTaskIds).Where(tagQuery).Find(&upTaskOutputs)
    for _, up := range t.Dag.Depandent.Up {
      for _, upTaskOutput := range upTaskOutputs {
        if up.Tag == upTaskOutput.Tag && up.UpTask == upTaskOutput.TaskName {
          input = append(
            input,
            Input{
              Value: upTaskOutput.Ret,
              Annotation: up.Annotation,
            },
          )
        }
      }
    }
  }
  if len(t.Dag.Depandent.Up) != len(input) {
    return input, &error.Error{
        Code: 31010,
        Hits: fmt.Sprintf(
            "JobID: %v; Group: %v; TaskName: %v; Dag.length==%v; database.length==%v",
            t.JobID,
            t.Group,
            t.Name,
            len(t.Dag.Depandent.Up),
            len(input),
        ),
    }
  }
  return input, nil
}


func SaveOutput(t model.Task, output []string) *error.Error {
  if len(t.Dag.Output) != len(output) {
    return &error.Error{
        Code: 31020,
        Hits: fmt.Sprintf(
            "JobID: %v; Group: %v; TaskName: %v; Dag.length==%v; process.length==%v",
            t.JobID,
            t.Group,
            t.Name,
            len(t.Dag.Output),
            len(output),
        ),
    }
  }

  var insertingTaskOutputs []model.TaskResult
  for index, tag := range t.Dag.Output {
    insertingTaskOutputs = append(insertingTaskOutputs, model.TaskResult{
      Task: t,
      Tag: tag,
      Ret: output[index],
      TaskName: t.Name,
    })
  }
  if len(insertingTaskOutputs) != 0 {
    db.DataBase.Debug().Create(insertingTaskOutputs)
  }
  return nil
}
