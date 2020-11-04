package main

import (
	"strings"
	"testing"
)

func TestObj(t *testing.T) {

	var v = `#EXT-X-KEY:METHOD=AES-128,URI="/20200803/2cglRcaJ/2000kb/hls/key.key"`
	start := strings.Index(v, "URI=\"") + 5
	end := strings.LastIndex(v, "\"")
	turl := v[start:end]
	end = start - 1 - 5
	start = strings.Index(v, "METHOD=") + 7
	method := v[start:end]
	println(turl, method)
}
