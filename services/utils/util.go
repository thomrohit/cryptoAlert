package utils

import (
	"log"
	"net/url"
	"strconv"
	"time"
)

const (
	DDMMYYY        = "02-01-2006"
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func Str2float(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal("conversion issue:", err)
	}
	return f
}

func Str2int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("conversion issue:", err)
	}
	return i
}

func Str2dateFormatStr(s string) string {
	t, _ := time.Parse(DDMMYYY, s)
	return t.Format(DDMMYYYYhhmmss)
}

func CreateUrl(params url.Values, newoffset int, API string) string {
	baseURL := "http://localhost:8080" + API
	v := url.Values{}
	v.Set("date", params.Get("date"))
	v.Set("offset", strconv.Itoa(newoffset))
	v.Set("limit", params.Get("limit"))

	perform := baseURL + "?" + v.Encode()
	return perform
}
