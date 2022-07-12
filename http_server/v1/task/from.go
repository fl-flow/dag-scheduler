package task


type TaskRunningForm struct {
  JobID     uint      `json:"job_id" binding:"required"`
  Group     string    `json:"group" binding:"required"`
  Task      string    `json:"task" binding:"required"`
}


type TaskCancelForm struct {
  JobID     uint      `json:"job_id" binding:"required"`
  Group     string    `json:"group" binding:"required"`
  Task      string    `json:"task" binding:"required"`
}
