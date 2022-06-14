package dagschedulerclient

import (
  "fmt"
	"log"
  "time"
  "bytes"
  "net/http"
	"io/ioutil"
)


func fetch(method string, uri string, jsonData []byte) []byte {
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
  if response.StatusCode != 200 {
    log.Fatalf("request for '%s' status : %v\n", url, response.StatusCode)
  }
	body, _ := ioutil.ReadAll(response.Body)
  return body
}
