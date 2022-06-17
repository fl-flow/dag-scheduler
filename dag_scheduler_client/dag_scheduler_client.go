package dagschedulerclient

import (
  "fmt"
)

var IP string
var Port string
var Protocol string

var Host string
var IPPort string


func init()  {
  // TODO:
  IP = "127.0.0.1"
  Port = "8000"
  Protocol = "http"
  IPPort = fmt.Sprintf("%s:%s", IP, Port)
  Host = fmt.Sprintf("%s://%s:%s", Protocol, IP, Port)
}
