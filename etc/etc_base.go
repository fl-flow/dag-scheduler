package etc


// scheduler server
var IsRunHttpApi bool = true
var SchedulerIp string = "127.0.0.1"
var SchedulerPort int = 8000


// scheduler loop
var IsRunSchedulerLoop bool = true
var SchedulerLoopMemory uint64 = 0

// distributed multiprocess
var MultiprocessIp string = "127.0.0.1"
var MultiprocessPort int = 8502

// resource coordinator
var ResourceCoordinatorIP string = "127.0.0.1"
var ResourceCoordinatorPort int = 8080


// cluster
var NodeId string = "http://127.0.0.1:8000" // TODO: check ip-port

// MYSQL
var MysqlUName string = "root"
var MysqlPWD string = ""
var MysqlHost string = "127.0.0.1"
var MysqlPort int = 3306
var MysqlDB string = "dag"
