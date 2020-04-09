package main

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func TestReadFile(t *testing.T) {
	fp := `/media/e/tdx/lday/sh603301.day`
	records, err := ReadFile(2020312, "sh603301", fp)
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
		err := ReadFolder(f, "http://127.0.0.1:8080/api/import", "000", 20200312)
		if err != nil {
			t.Error(err)
			debug.PrintStack()
		}
	}
}
