package main

import (
	"github.com/robfig/cron/v3"
	"time"
)

func main () {
	c := cron.New()
	c.AddFunc("@every 5s",func() {
		fileName := time.Now().Format("20060102150405")+".json"
		getTransactionsByAccount(fileName)
	})
	c.Start()
	for {
		time.Sleep(time.Second)
	}
}