package error


var Conf = map[int]string {
  0: "success",

  10000: "parser error",

  11000: "dag parser error", // base
  11010: "dag parser error( required task.tag; task and tag can't contain '.' )",
  11020: "dag parser error( task.tag; task not exits)",
  11021: "dag parser error( some dag's group don't found in parameter's group)",
  11030: "dag parser error( loop found)",
  11040: "dag parser error( cmd is required )",

  12000: "parameter parser error", // base
  12010: "num of parameter is not equal to num of dag",
  12020: "num of parameter is not equal to num of dag",


  20000: "loop error",

  21000: "init task error",

  22000: "ready task error",

  30000: "ready task runner error",
  31010: "runner error ( some dag input is not existed in database)",
  31020: "runner error ( dag out.length is not equal to process's out.length )",


  110000: "job http api error", // base
  110010: "job http api error (no tasks)",


  80000: "client error",
  80010: "client remote error",
}
