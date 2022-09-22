package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
	"github.com/langwan/langgo/core/log"
	"sync"
)

func worker4() {
	id := 1
	n := 100
	langgo.Run(&mysql.Instance{})
	mysql.Get("main").AutoMigrate(&Account{}, &Account2{})
	mysql.Main().Delete(&Account{}, "id=?", id)

	mysql.Main().Create(&Account{
		Id:      id,
		Name:    "chi",
		Balance: 0,
	})
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			acc := Account{}
			res := mysql.Main().Find(&acc, "id=?", id)
			if res.Error != nil {
				log.Logger("app", "worker4").Warn().Err(res.Error).Send()
			}
			acc.Balance += 1
			res = mysql.Main().Save(&acc)
			if res.Error != nil {
				log.Logger("app", "worker4").Warn().Err(res.Error).Send()
			}

		}()
	}
	wg.Wait()
	log.Logger("app", "worker4").Info().Msg("ok")
}
