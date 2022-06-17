package dagschedulerclient

import (
  "fmt"
	"log"
  "time"
  "bytes"
  "net/http"
	"io/ioutil"

  "github.com/fl-flow/dag-scheduler/common/error"
)


func fetch(method string, uri string, jsonData []byte) ([]byte, *error.Error) {
  url := fmt.Sprintf("%s%s", Host, uri)
  request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("new request to '%s' failed: %v\n", url, err)
	}
  var client = &http.Client{
		Timeout:   time.Second * 30,
	}
  response, err := client.Do(request)
  if err != nil {
		log.Fatalf("request for '%s' failed: %v\n", url, err)
	}
  defer response.Body.Close()
  body, _ := ioutil.ReadAll(response.Body)
  if response.StatusCode != 200 {
    log.Println("request for '%s' status : %v\n body: %v\n", url, response.StatusCode, string(body))
    return body, &error.Error{
      Code: 80010,
      Hits: string(body),
    }
  }
  return body, nil
}
