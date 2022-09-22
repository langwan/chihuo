package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
	"github.com/langwan/langgo/core/log"
)

func GetAccount() (acc *Account, err error) {
	acc = &Account{}
	res := mysql.Main().First(acc)
	if res.Error != nil {
		log.Logger("app", "model").Err(err).Send()
		return nil, res.Error
	}
	return acc, nil
}

func worker3() {
	langgo.Run(&mysql.Instance{})

	mysql.Main().Delete(&Account{}, "id=?", 1)

	mysql.Main().Create(&Account{
		Id:   1,
		Name: "chi",
	})

	acc, err := GetAccount()
	if err != nil {
		log.Logger("app", "worker3").Info().Err(err).Interface("acc", acc).Send()

	} else {
		log.Logger("app", "worker3").Info().Interface("acc", acc).Msg("ok")
	}

	mysql.Main().Delete(&Account{}, "id=?", 1)
	acc, err = GetAccount()

	if err != nil {
		log.Logger("app", "worker3").Info().Err(err).Interface("acc", acc).Send()

	} else {
		log.Logger("app", "worker3").Info().Interface("acc", acc).Msg("ok")
	}
}
