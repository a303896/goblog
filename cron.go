package main

import (
	"blog/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("cron is starting...")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("models.CleanAllTag is running...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("models.CleanAllArticle is running...")
		models.ClearAllArticle()
	})
	c.Start()
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
			case <-t1.C:
				t1.Reset(time.Second * 10)
		}
	}
}
