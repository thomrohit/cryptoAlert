package alert

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/thomrohit/cryptoAlert/services/email"
	"github.com/thomrohit/cryptoAlert/services/utils"
)

type Crypto struct {
	Bitcoin CoinInfo
}

type CoinInfo struct {
	Usd float64
}

func CryptoAlert() {
	log.Info("Start crypto alert service")
	c := new(Crypto)
	minPrice := utils.Str2float(os.Getenv("MIN"))
	maxPrice := utils.Str2float(os.Getenv("MAX"))

	getJson("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd", c)

	DBinsert(c.Bitcoin.Usd)
	if c.Bitcoin.Usd <= minPrice {
		log.Info("less than min price ")
		email.SendEmail("Alert!!! Below Min Price, Price has changed to " + strconv.FormatFloat(c.Bitcoin.Usd, 'g', -1, 64))
	} else if c.Bitcoin.Usd >= maxPrice {
		log.Info("more than max price ")
		email.SendEmail("Alert!!! Exceed Max Price, Price has changed to " + strconv.FormatFloat(c.Bitcoin.Usd, 'g', -1, 64))
	}
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Error("Failed to get Request", err)
		return err
	}
	defer r.Body.Close()
	log.Info("Convert request to JSON")
	return json.NewDecoder(r.Body).Decode(target)
}
