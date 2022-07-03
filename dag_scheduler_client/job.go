package dagschedulerclient

import (
  "log"
  "encoding/json"
  "github.com/mitchellh/mapstructure"

  "github.com/fl-flow/dag-scheduler/common/error"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


func CreateJob(data interface{}) (model.Job, *error.Error) {
  b, err := json.Marshal(data)
  if err != nil {
    log.Fatalf("data json dumps error:  %v\n", err)
  }
  data, e := fetch("POST", "/api/v1/job/", b)
  var job model.Job
  if e != nil {
    return job, e
  }
  mapstructure.Decode(data, &job)
  return job, nil
}


func ListJob(size int, page int) ([]model.Job, *error.Error) {
  var jobs []model.Job
  body, e := fetch("GET", "/api/v1/job/", []byte(`{}`))
  if e != nil {
    return jobs, e
  }
  var lr ListRet
  err_ := json.Unmarshal([]byte(body), &lr)
  if err_ != nil {
    log.Fatalf("data json loads error:  %v\n", err_)
  }
  mapstructure.Decode(lr.Data.List, &jobs)
  return jobs, nil
}
