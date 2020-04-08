package main

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	fp := `/media/e/tdx/lday/sh603301.day`
	records, err := ReadFile(2020312, fp)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		t.Error(err)
	}
	for _, e := range records {
		t.Log(fmt.Sprintf("%+v", e))
	}
}
func TestReadFolder(t *testing.T) {
	fds := []string{"/media/e/tdx/lday", "/media/e/tdx/ll/lday"}
	for _, f := range fds {
		ReadFolder(f, "http://127.0.0.1:8080", "000", 2020312)
	}
}
