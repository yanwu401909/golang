package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFatchData(t *testing.T) {
	url := "https://api.ip.sb/geoip/"
	data, err := fatchData(url)
	assert.Empty(t, err)
	// t.Log(err)
	t.Logf("Data: %s ", data)
}
