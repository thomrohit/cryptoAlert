package prices

import (
	"encoding/json"
	"net/http"

	"github.com/thomrohit/cryptoAlert/services/alert"
	"github.com/thomrohit/cryptoAlert/services/utils"
)

const (
	API = "/api/prices/btc"
)

func PricePage(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	offset := utils.Str2int(params.Get("offset"))
	prices := alert.GetPrice(utils.Str2dateFormatStr(params.Get("date")), utils.Str2int(params.Get("limit")), offset)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CryptoResponse{
		Url:  utils.CreateUrl(params, offset, API),
		Next: utils.CreateUrl(params, offset+offset, API),
		Data: prices})
}

type CryptoResponse struct {
	Url  string         `json:"url"`
	Next string         `json:"next"`
	Data []*alert.Price `json:"data"`
}
