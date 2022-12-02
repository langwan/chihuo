package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("30 * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})
}
