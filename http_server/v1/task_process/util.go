package taskprocess

import (
  "log"
  "bytes"
  "bufio"
)


type ProcessRetChanDataType struct {
  Data      []byte
  Done      bool
}


type BlockHeader struct{
  ContentDisposition  string
  Name                string
  ContentType         string
}


const (
  CONTENT_DISPOSITION = ("Content-Disposition: ")
  NAME =  ("name=\"")
  FILENAME = ("filename=\"")
  CONTENT_TYPE = ("Content-Type: ")
)

const BOUNDARY string = "; boundary="


// TODO: // BUG: read content-length
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
      log.Fatalf("error upload")
    }
  }
  return b, false
}


func readData(reader *bufio.Reader, boundary []byte, ch chan ProcessRetChanDataType, h BlockHeader) {
  former, prefix, err := reader.ReadLine()
  if err != nil || prefix || bytes.Index(former, boundary) > 0 {
    ch <- ProcessRetChanDataType {
      Done: true,
    }
    return
  }
  for true {
    part, prefix, err := reader.ReadLine()
    if err != nil || prefix || bytes.Index(part, boundary) > 0 {
      ch <- ProcessRetChanDataType {
        Data: former,
        Done: false,
      }
      ch <- ProcessRetChanDataType {
        Done: true,
      }
      return
    }
    ch <- ProcessRetChanDataType {
      Data: bytes.Join([][]byte{former, []byte("\n")}, []byte("")),
      Done: false,
    }
    former = part
  }
}
