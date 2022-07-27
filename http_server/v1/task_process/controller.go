package taskprocess

import (
  "log"
  "bufio"
  "strings"
)


func TaskProcessController(
  chRes chan ProcessRetChanDataType,
  reader *bufio.Reader,
  boundary []byte,
) {
  readToBoundary(reader, boundary)
  var cmd string
  for true {
    b, cont := parseBlockHeader(reader)
    if !cont {
      break
    }
    var ch chan ProcessRetChanDataType
    if b.Name == "cmd" {
      ch = make(chan ProcessRetChanDataType)
      go func () {
        for true {
          d := <-ch
          if d.Done {
            return
          }
          cmd = strings.Trim(string(d.Data), "\n")
          log.Println("got Cmd: ", cmd)
        }
      }()
      readData(reader, boundary, ch, b)
    } else if b.Name == "data" {
      ch = make(chan ProcessRetChanDataType, 128)
      if cmd == "" {
        chRes <- ProcessRetChanDataType{
          Data: []byte("error no cmd"),
          Done: true,
        }
        return
      }
      go readData(reader, boundary, ch, b)
      RunProcess(cmd, ch, chRes)
      return
    } else {
      readToBoundary(reader, boundary)
      continue
    }
  }
  if cmd == "" {
    chRes <- ProcessRetChanDataType{
      Data: []byte("error no cmd"),
      Done: true,
    }
  }
}
