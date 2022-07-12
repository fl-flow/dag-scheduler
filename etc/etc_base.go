package etc


// scheduler server
var IsRunHttpApi bool = true
var SchedulerIp string = "127.0.0.1"
var SchedulerPort int = 8000


// scheduler loop
var IsRunSchedulerLoop bool = true
var SchedulerLoopMemory uint64 = 0


// resource coordinator
var ResourceCoordinatorIP string = "127.0.0.1"
var ResourceCoordinatorPort int = 8080


// cluster
var NodeId string = "127.0.0.1:8000" // TODO: check ip-port
