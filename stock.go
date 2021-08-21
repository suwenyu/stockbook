package stockbook

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/suwenyu/stockbook/utils"
)

type tsMapData map[int64][][]string

type stockInfo struct {
	open        float64
	high        float64
	low         float64
	price       float64
	change      float64
	dailychange float64
	volume      float64
	date        time.Time
}

// Data start with stock no, date.
type Data struct {
	No        string
	Name      string
	Date      time.Time
	RawData   [][]string
	TsMapData tsMapData
	Exchange  string
	StockInfo []stockInfo
}

// URL return stock csv url path.
func (d *Data) URL() string {
	switch d.Exchange {
	case "twse":
		return fmt.Sprintf("%s%s", utils.TWSEURL,
			fmt.Sprintf(utils.TWSESTOCK, d.Date.Year(), d.Date.Month(), d.Date.Day(), d.No))
	}

	return ""
}

func NewTWSE(No string, Date time.Time) *Data {
	return &Data{
		No:        No,
		Date:      Date,
		Exchange:  "twse",
		TsMapData: make(tsMapData),
	}
}

type StockPayload struct {
	Stat   string     `json:"stat"`
	Date   string     `json:"date"`
	Title  string     `json:"title"`
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
	Notes  []string   `json:"notes"`
}

func (d *Data) Get() string {
	var msg StockPayload

	json_bytes, err := utils.GetJson(d.URL())
	if err != nil {
		return ""
	}

	err = json.Unmarshal(json_bytes, &msg)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(utils.PrettyPrint(msg))
	return ""
}
