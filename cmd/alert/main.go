package main

import (
	"net/http"

	"github.com/jasonlvhit/gocron"
	"github.com/labstack/gommon/log"
	"github.com/thomrohit/cryptoAlert/services/alert"
	"github.com/thomrohit/cryptoAlert/services/prices"
)

func main() {
	alert.ConnectDB()
	gocron.Start()
	gocron.Every(30).Seconds().From(gocron.NextTick()).Do(alert.CryptoAlert)
	http.HandleFunc(prices.API, prices.PricePage)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
