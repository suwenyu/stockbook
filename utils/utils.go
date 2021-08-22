package utils

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
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

var dateReg = regexp.MustCompile(`([\d]{2,})/([\d]{1,2})/([\d]{1,2})`)

func ParseDate(strDate string) (int64, error) {
	p := dateReg.FindStringSubmatch(strDate)
	if len(p) == 0 {
		err := errors.New("Date format is not legal.")
		return 0, err
	}
	year, _ := strconv.Atoi(p[1])
	mon, _ := strconv.Atoi(p[2])
	day, _ := strconv.Atoi(p[3])
	return time.Date(year+1911, time.Month(mon), day, 0, 0, 0, 0, TaipeiTimeZone).Unix(), nil
}
