package form


type TaskRunningForm struct {
  JobID     uint      `json:"job_id"`
  Group     string    `json:"group"`
  Task      string    `json:"task"`
}
