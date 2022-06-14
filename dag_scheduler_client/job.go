package dagschedulerclient

import (
  "log"
  "encoding/json"
  "dag/common/error"
  "dag/common/db/model"
  "github.com/mitchellh/mapstructure"
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
  body := fetch("POST", "/v1/job/", b)
  err_ := json.Unmarshal([]byte(body), &ret)
  if err_ != nil {
    log.Fatalf("data json loads error:  %v\n", err_)
  }
  var job model.Job
  mapstructure.Decode(ret.Data, &job)
  return job, nil
}
