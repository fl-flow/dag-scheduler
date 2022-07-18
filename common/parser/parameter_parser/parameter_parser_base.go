package parameterparser


type Resource struct {
  Memory    uint64   `json:"memory"`
}


type Setting struct {
  Resource  Resource  `json:"resource"`
}


type TaskParameter struct {
  Args      map[string]interface{}  `json:"args"`
  Setting   Setting                 `json:"setting"`
}


type Parameter struct {
  Common      string                      `json:"common"`
  Tasks       map[string]TaskParameter    `json:"tasks"`
}
