package stockbook

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/suwenyu/stockbook/utils"
)

var GetJson = utils.GetJson

type tsMapData map[int64][][]string

// FmtData is struct for daily data format.
type StockInfo struct {
	Date       time.Time
	Volume     uint64  //成交股數
	TotalPrice uint64  //成交金額
	Open       float64 //開盤價
	High       float64 //最高價
	Low        float64 //最低價
	Price      float64 //收盤價
	Range      float64 //漲跌價差
	Totalsale  uint64  //成交筆數
}

// Data start with stock no, date.
type Data struct {
	No           string
	Name         string
	Date         time.Time
	DailyMapData map[int64]interface{}
	MonthMapData tsMapData
	Exchange     string
	StockInfo    []StockInfo
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
		No:           No,
		Date:         Date,
		Exchange:     "twse",
		MonthMapData: make(tsMapData),
		DailyMapData: make(map[int64]interface{}),
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

func (d *Data) Get() ([][]string, error) {
	var msg StockPayload
	monthDateUnix := time.Date(d.Date.Year(), d.Date.Month(), 1, 0, 0, 0, 0, d.Date.Location()).Unix()

	json_bytes, err := GetJson(d.URL())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(json_bytes, &msg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	d.MonthMapData[monthDateUnix] = msg.Data
	d.GetDailyMapData(msg.Data)

	return d.MonthMapData[monthDateUnix], nil
}

// MinusMonth would sub one month.
func (d *Data) MinusMonth() {
	year, month, _ := d.Date.Date()
	d.Date = time.Date(year, month-1, 1, 0, 0, 0, 0, d.Date.Location())
}

// RetrievePrevMonth would do Round() and Get() for input months
func (d *Data) RetrievePrevMonth(month int) {
	d.Get()
	for i := 0; i < month; i++ {
		d.MinusMonth()
		d.Get()
	}
}

func (d *Data) GetDailyMapData(data [][]string) error {

	for _, v := range data {
		datets, err := utils.ParseDate(v[0])
		if err != nil {
			return err
		}
		d.DailyMapData[datets] = v
	}

	return nil
}

func (d Data) SortedKeys(m map[int64]interface{}) []int64 {

	keys := make([]int64, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })
	return keys
}

// FormatData format daily data.
func (d *Data) FormatData() []StockInfo {
	var (
		data   StockInfo
		result []StockInfo
	)

	sortedKeyList := d.SortedKeys(d.DailyMapData)

	result = make([]StockInfo, len(sortedKeyList))
	op := make([]string, 0)

	for i, ts := range sortedKeyList {
		op = d.DailyMapData[ts].([]string)

		vdate, _ := utils.ParseDate(op[0])
		data.Date = time.Unix(vdate, 0).In(utils.TaipeiTimeZone)
		data.Volume, _ = strconv.ParseUint(strings.Replace(op[1], ",", "", -1), 10, 64)
		data.Open, _ = strconv.ParseFloat(op[3], 64)
		data.Low, _ = strconv.ParseFloat(op[5], 64)
		data.Price, _ = strconv.ParseFloat(op[6], 64)
		data.Range, _ = strconv.ParseFloat(op[7], 64)
		data.Totalsale, _ = strconv.ParseUint(strings.Replace(op[8], ",", "", -1), 10, 64)
		result[i] = data
	}

	return result
}
