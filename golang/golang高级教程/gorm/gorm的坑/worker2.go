package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
)

func worker2() {

	langgo.Run(&mysql.Instance{})

	mysql.Get("main").AutoMigrate(&Account{}, &Account2{})

	a1 := Account{Name: "chihuo"}
	a2 := Account2{Name: "famingjia"}

	mysql.Main().Create(&a1)
	mysql.Main().Create(&a2)

	mysql.Main().Delete(&a1)

	mysql.Main().Delete(&a2)
	//mysql.Main().Unscoped().Delete(&a2)
}
