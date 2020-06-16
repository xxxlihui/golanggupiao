package main

import (
	"net/url"
	"testing"
)

func TestObj(t *testing.T) {
	u, _ := url.Parse("http://www.baidu.com/s/s/s")
	u2, _ := u.Parse("/a/a/a")
	t.Log(u2.String())
}
