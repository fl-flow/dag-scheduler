package taskprocess

import (
  "log"
  "strings"
  "bytes"
  "fmt"
  "bufio"

  "github.com/gin-gonic/gin"
)


type BlockHeader struct{
  ContentDisposition  string
  Name                string
  ContentType         string
  ContentLength       int64
}

const (
  CONTENT_DISPOSITION = ("Content-Disposition: ")
  NAME =  ("name=\"")
  FILENAME = ("filename=\"")
  CONTENT_TYPE = ("Content-Type: ")
)

func readToBoundary(reader *bufio.Reader, boundary []byte) {
  for true {
    part, prefix, err := reader.ReadLine()
    if err != nil {
      return
    }
    if prefix {
      return
    }
    if loc := bytes.Index(part, boundary); loc > 0 {
      break
    }
  }
}

func parseBlockHeader(reader *bufio.Reader) (BlockHeader, bool) {
  cont := true
  var b BlockHeader
  for cont {
    part, prefix, err := reader.ReadLine()
    if err != nil {
      return b, false
    }
    if prefix {
      return b, false
    }
    if string(part) == "" {
      return b, true
    }
    if bytes.HasPrefix(part, []byte(CONTENT_DISPOSITION)) {
      arr1 := bytes.Split(part[len(CONTENT_DISPOSITION):], []byte("; "))
      b.ContentDisposition = string(arr1[0])
      if bytes.HasPrefix(arr1[1], []byte(NAME)){
        b.Name = string(arr1[1][len(NAME):len(arr1[1])-1])
      }
    }else if bytes.HasPrefix(part, []byte(CONTENT_TYPE)) {
      b.ContentType = string(part[len(CONTENT_TYPE):])
    } else {
      fmt.Println(string(part))
      log.Fatalf("error upload")
    }
  }
  return b, false
}


func readData(reader *bufio.Reader, boundary []byte, ch chan []byte) {
  for true {
    part, prefix, err := reader.ReadLine()
    if err != nil {
      return
    }
    if prefix {
      return
    }
    if loc := bytes.Index(part, boundary); loc > 0 {
      break
    }
    ch <- part
  }
}


func TaskProcess(c *gin.Context) {
  var content_length int64
	content_length = c.Request.ContentLength
	if content_length<=0 || content_length > 1024*1024*1024*2 {
		log.Printf("content_length error\n")
    // TODO:
		return
	}
	contentType_, has_key := c.Request.Header["Content-Type"]
	if !has_key{
		log.Printf("Content-Type error\n")
    // TODO:
		return
	}
	if len(contentType_) != 1{
		log.Printf("Content-Type count error\n")
    // TODO:
		return
	}
	contentType := contentType_[0]
	const BOUNDARY string = "; boundary="
	loc := strings.Index(contentType, BOUNDARY)
	if -1 == loc{
		log.Printf("Content-Type error, no boundary\n")
    // TODO:
		return
	}
	boundary := []byte(contentType[(loc+len(BOUNDARY)):])
	log.Printf("[%s]\n\n", boundary)
  var cmd string
  reader := bufio.NewReader(c.Request.Body)
  readToBoundary(reader, boundary)
  for true {
    b, cont := parseBlockHeader(reader)
    if !cont {
      break
    }
    var ch chan []byte
    if b.Name == "cmd" {
      ch = make(chan []byte)
      go func (ch chan []byte) {
        cmd = string(<-ch)
        fmt.Println(cmd, "// TO Cmd: ")
      }(ch)
    } else if b.Name == "data" {
      ch = make(chan []byte)
      // todo assert cmd is exsited
      go func (ch chan []byte) {
        for true {
          fmt.Println(string(<-ch), "// TO input data: ")
        }
      }(ch)
    } else {
      readToBoundary(reader, boundary)
      continue
    }
    readData(reader, boundary, ch)
  }
  if cmd == "" {
    fmt.Println("error no cmd")
  }
  fmt.Println(cmd)
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("%s", "ok"),
	})
}
