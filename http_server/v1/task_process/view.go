package taskprocess

import (
  "io"
  "bufio"
  "strings"

  "github.com/gin-gonic/gin"
  "github.com/fl-flow/dag-scheduler/http_server/http/response"
)


func TaskProcess(c *gin.Context) {
  var content_length int64
	content_length = c.Request.ContentLength
	if content_length<=0 || content_length > 1024*1024*1024*2 {
    response.R(
      c,
      100,
      "content_length error",
      "content_length error",
    )
		return
	}
	contentType_, has_key := c.Request.Header["Content-Type"]
	if !has_key{
    response.R(
      c,
      100,
      "Content-Type error",
      "Content-Type error",
    )
		return
	}
	if len(contentType_) != 1{
    response.R(
      c,
      100,
      "Content-Type count error",
      "Content-Type count error",
    )
		return
	}
	contentType := contentType_[0]
	loc := strings.Index(contentType, BOUNDARY)
	if -1 == loc {
    response.R(
      c,
      100,
      "Content-Type error, no boundary",
      "Content-Type error, no boundary",
    )
		return
	}
	boundary := []byte(contentType[(loc+len(BOUNDARY)):])
  reader := bufio.NewReader(c.Request.Body)
  chRes := make(chan ProcessRetChanDataType)

  // TODO: error controller
  go TaskProcessController(chRes, reader, boundary)
  w := c.Writer
  header := w.Header()
  header.Set("Transfer-Encoding", "chunked")
  header.Set("Content-Type", "text/event-stream")
  c.Stream(func(w io.Writer) bool {
      d := <- chRes
      if d.Done {
        if len(d.Data) != 0 {
          w.Write(d.Data)
        }
        return false
      }
      w.Write(d.Data)
      return true
  })
}
