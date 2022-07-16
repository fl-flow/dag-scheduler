package db

import (
  "log"
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"

  "github.com/fl-flow/dag-scheduler/etc"
  "github.com/fl-flow/dag-scheduler/common/db/model"
)


func init()  {
  db, er := gorm.Open(
    mysql.Open(
      fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
        etc.MysqlUName,
        etc.MysqlPWD,
        etc.MysqlHost,
        etc.MysqlPort,
        etc.MysqlDB,
        "10s",
      ),
    ),
    &gorm.Config{
      DisableForeignKeyConstraintWhenMigrating: true,
    },
  )
  if er != nil {
    log.Fatalln("error db connect")
  }

  DataBase = db
  db.AutoMigrate(
    &model.TaskLink{},
  )
  db.AutoMigrate(
    &model.Job{},
    &model.Task{},
    &model.TaskResult{},
  )
}
