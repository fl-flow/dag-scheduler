package dagschedulerclient


type ListData struct {
  Count     int           `json:"count"`
  List      []interface{} `json:"list"`
  Page      int           `json:"page"`
  Size      int           `json:"size"`
}


type ListRet struct {
  Code    int       `json:"code"`
  Data    ListData  `json:"data"`
  Msg     string    `json:"msg"`
}
