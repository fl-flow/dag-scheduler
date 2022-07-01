package dagparser


// raw
type DagTask struct {
  Input       []string  `json:"input"`
  Output      []string  `json:"output"`
  Cmd         []string   `json:"cmd"`
  ValidateCmd string    `json:"validate_cmd"`
}


// parsered
type DagTaskMap map[string]DagTask


type TaskInput struct {
  UpTask        string
  Tag           string
  Annotation    string
}


type TaskDepandent struct {
  Up          []TaskInput
  Down        []string
}

type TasksDepandentMap map[string]*TaskDepandent

type TaskParsered struct {
  Name          string
  Depandent     TaskDepandent
  Output        []string
  Cmd           []string
  ValidateCmd   string
}

type TaskParseredList []TaskParsered
