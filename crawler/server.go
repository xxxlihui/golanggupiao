package crawler

import (
	"errors"
	"fmt"
	"nn/data"
	"runtime/debug"
)

func GetByCodes(codes []string) (records []*data.DayRecord, err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v", p)
			debug.PrintStack()
			err = errors.New(fmt.Sprintf("panic recover! p: %v", p))
		}
	}()
	return sinaGetByCodes(codes)
}
