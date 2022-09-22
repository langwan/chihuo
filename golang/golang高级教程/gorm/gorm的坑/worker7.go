package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
	"github.com/langwan/langgo/core/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

func worker7() {
	id := 1
	langgo.Run(&mysql.Instance{})
	mysql.Get("main").AutoMigrate(&Account{}, &Account2{})
	mysql.Main().Delete(&Account{}, "id=?", id)

	mysql.Main().Create(&Account{
		Id:      id,
		Name:    "chi",
		Balance: 0,
	})
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		j := i
		go func() {
			defer func() {
				if err := recover(); err != nil {
				}
				wg.Done()
			}()
			mysql.Main().Transaction(func(tx *gorm.DB) error {

				acc := Account{}
				res := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&acc, "id=?", id)
				if res.Error != nil {
					log.Logger("app", "worker7").Warn().Err(res.Error).Send()
					return res.Error
				}
				acc.Balance += 1
				res = tx.Save(&acc)
				if res.Error != nil {
					log.Logger("app", "worker7").Warn().Err(res.Error).Send()
					return res.Error
				}
				if j == 2 || j == 4 {
					panic("j = 2 or 4 exit")
				}
				return nil
			})

		}()
	}
	wg.Wait()
	log.Logger("app", "worker4").Info().Msg("ok")
}
