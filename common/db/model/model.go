package model

import (
  "time"
)


type BaseModel struct {
  ID        uint        `gorm:"primarykey"`
  CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}


type Job struct {
  BaseModel
  Status          JobStatus     `gorm:"type:int"`
  Name            string
  Description     string
  Dag             JobDag        `gorm:"type:json"`
  RawDag          JobRawDagmap  `gorm:"type:json"`
  Parameter       JobParameter  `gorm:"type:json"`
  Tasks           []Task
}


type Task struct {
  BaseModel
  JobId           uint
  Job             Job           `gorm:"ForeignKey:JobId;AssociationForeignKey:ID"`
  Group           string
  Status          TaskStatus    `gorm:"type:int"`
  Name            string
  Description     string
  Pid             int
  OrderInJob      int
  UpTasks         TaskUpTasks    `gorm:"type:json"`
  MemoryLimited   int
  Cmd             string
  ValidateCmd     string
}


type TaskResult struct {
  BaseModel
  TaskId          uint
  Task            Task          `gorm:"ForeignKey:TaskId;AssociationForeignKey:ID"`
  Tag             string
  Ret             string
}
