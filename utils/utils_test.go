package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestTimeData struct {
	input  string
	expect int64
}

func TestParseDate(t *testing.T) {
	TestDataList := []TestTimeData{
		{"110/08/23", time.Date(2021, 8, 23, 0, 0, 0, 0, TaipeiTimeZone).Unix()},
		{"110/02/29", time.Date(2021, 3, 1, 0, 0, 0, 0, TaipeiTimeZone).Unix()},
		{"110/06/31", time.Date(2021, 7, 1, 0, 0, 0, 0, TaipeiTimeZone).Unix()},
		{"110/06/31", time.Date(2021, 7, 1, 0, 0, 0, 0, TaipeiTimeZone).Unix()},
		{" ", time.Now().Unix()},
		{"", time.Now().Unix()},
	}

	for _, data := range TestDataList {
		ts, err := ParseDate(data.input)
		if err == nil {
			assert.Nil(t, err)
			assert.Equal(t, ts, data.expect)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestPrettyPrint(t *testing.T) {
	var i interface{}
	i = 12
	output := PrettyPrint(i)
	assert.Equal(t, output, "12")
}
