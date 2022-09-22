package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
	"github.com/langwan/langgo/core/log"
	"gorm.io/gorm"
)

type Account struct {
	Id      int `gorm:"primaryKey;autoIncrement"`
	Name    string
	Balance int
}

type Account2 struct {
	gorm.Model
	Name    string
	Balance int
}

func worker1() {
	langgo.Run(&mysql.Instance{})
	var one int
	mysql.Main().Raw("SELECT11").Scan(&one)
	log.Logger("app", "worker1").Info().Int("one", one).Send()
}
