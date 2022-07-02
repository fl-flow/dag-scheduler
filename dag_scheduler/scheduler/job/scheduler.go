package jobscheduler

import (
  "sync"
)


var Wait *sync.WaitGroup
var Limit int = 5000


func init()  {}
