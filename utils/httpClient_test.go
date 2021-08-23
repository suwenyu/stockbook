package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJson(t *testing.T) {
	resp, _ := GetJson("https://www.google.com")
	assert.NotNil(t, resp)

	_, err := GetJson("https://fake.url.com")
	assert.NotNil(t, err)
}
