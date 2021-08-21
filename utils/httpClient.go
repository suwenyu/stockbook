package utils

import (
	"io/ioutil"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string) ([]byte, error) {
	r, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {

		return nil, err
	}
	defer r.Body.Close()

	return b, nil
}
