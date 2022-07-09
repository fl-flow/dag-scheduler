package model

import (
  "time"
)


type BaseModel struct {
  ID        uint        `gorm:"primarykey;type:bigint auto_increment"`
  CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}


type Job struct {
  BaseModel
  ID              uint          `gorm:"primarykey;type:bigint auto_increment"`
  Status          JobStatus     `gorm:"type:int"`
  Name            string
  Description     string
  Dag             JobDag        `gorm:"type:json"`
  RawDag          JobRawDagmap  `gorm:"type:json"`
  Parameter       JobParameter  `gorm:"type:json"`
  Tasks           []Task
  NotifyUrl       string
}


type Task struct {
  BaseModel
  ID              uint          `gorm:"primarykey;type:bigint auto_increment"`
  JobID           uint          `gorm:"type:bigint"`
  Job             Job
  Group           string
  Status          TaskStatus    `gorm:"type:int"`
  Name            string
  Description     string
  Pid             int
  OrderInJob      int
  Dag             TaskDag       `gorm:"type:json"`
  Parameters      TaskParameter `gorm:"type:json"`
  CommonParameter string
  UpTasks         []Task        `gorm:"many2many:TaskLink;joinForeignKey:DownId;joinReferences:UpId"`
  DownTasks       []Task        `gorm:"many2many:TaskLink;joinForeignKey:UpId;joinReferences:DownId"`
  // MemoryLimited   uint64
  // Cmd             CmdType       `gorm:"type:json"`
  // ValidateCmd     string
  CmdRet          string
  NotifyUrl       string
  GotCmdToRun     bool
  RunOnNode       string
}


type TaskLink struct {
	UpId           uint            `gorm:"type:bigint"`
	DownId         uint            `gorm:"type:bigint"`
}


type TaskResult struct {
  BaseModel
  TaskId          uint            `gorm:"type:bigint"`
  Task            Task
  Tag             string
  TaskName        string
  Ret             string
}
