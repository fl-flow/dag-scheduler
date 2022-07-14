package federation

import (
	"log"
	"fmt"
  "time"
  "bytes"
  "net/http"
  "encoding/json"
)


func (n Node) Cancel (
  jobId uint,
  group string,
  taskName string,
) {
  b, _ := json.Marshal(map[string]interface{} {
    "job_id": jobId,
    "group": group,
    "task": taskName,
  })
	url := fmt.Sprintf("%v/api/v1/task/cancel/", n.ID)
  request, _ := http.NewRequest(
    "POST",
    url,
    bytes.NewBuffer([]byte(b)),
  )
  _, err := (&http.Client{
		Timeout:   time.Second * 30,
	}).Do(request)
  if err != nil {
    // TODO: 重发
		log.Fatalf("request for '%s' failed: %v\n", url, err)
	}
  // TODO: ?????
}
