package dagschedulerclient

import (
  "log"
  "encoding/json"
  "github.com/mitchellh/mapstructure"

  "github.com/fl-flow/dag-scheduler/common/error"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


type CreateJobRet struct {
  Code    int             `json:"code"`
  Data    interface{}     `json:"data"`
  Msg     interface{}     `json:"msg"`
}


func CreateJob(data interface{}) (model.Job, *error.Error) {
  var ret CreateJobRet
  b, err := json.Marshal(data)
  if err != nil {
    log.Fatalf("data json dumps error:  %v\n", err)
  }
  body, e := fetch("POST", "/api/v1/job/", b)
  var job model.Job
  if e != nil {
    return job, e
  }
  err_ := json.Unmarshal([]byte(body), &ret)
  if err_ != nil {
    log.Fatalf("data json loads error:  %v\n", err_)
  }
  mapstructure.Decode(ret.Data, &job)
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
