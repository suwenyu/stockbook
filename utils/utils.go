package utils

import (
	"encoding/json"
	"time"
)

// TaipeiTimeZone is for time.Data() setting.
var TaipeiTimeZone = time.FixedZone("Asia/Taipei", 8*3600)

// TWSE base url.
const (
	TWSEURL   string = "http://www.twse.com.tw"
	TWSESTOCK string = "/exchangeReport/STOCK_DAY?response=json&date=%d%02d%02d&stockNo=%s"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
