package main

import (
	"github.com/HugoJBello/personal-monitor-golang/sources"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var trelloSource sources.TrelloSource

func init() {
	godotenv.Load("./.env")

	key := os.Getenv("TRELLOKEY")
	token := os.Getenv("TRELLOTOKEN")
	list := os.Getenv("TRELLODEFAULTLIST")

	trelloSource = sources.TrelloSource{
		TrelloKey:   key,
		TrelloToken: token,
		DefaultList: list,
	}
}

func extractFromAllSources() {
	trelloResult,_ := trelloSource.Extract()

	_ = trelloSource.Display(trelloResult)
}


func main() {

	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				extractFromAllSources()
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(time.Second * 1000000)
}
