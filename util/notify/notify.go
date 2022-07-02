package notify

import (
	"log"
  "time"
  "bytes"
  "net/http"
  "encoding/json"
)


func NotifyStatus(url string, status int, type_ string, id uint, extra interface{}) {
  b, _ := json.Marshal(map[string]interface{} {
    "status": status,
    "type": type_,
    "id": id,
		"extra": extra,
  })
  request, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(b)))
  _, err := (&http.Client{
		Timeout:   time.Second * 30,
	}).Do(request)
  if err != nil {
    // TODO: 重发
		log.Fatalf("request for '%s' failed: %v\n", url, err)
	}
}
