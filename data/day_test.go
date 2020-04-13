package data

import (
	"testing"
	"unsafe"
)

func TestStruct(t *testing.T) {
	day := &DayRecord{
		Day:      0,
		Open:     0,
		Code:     "12345612",
		Name:     "发放发生",
		High:     0,
		Low:      0,
		Close:    0,
		PreClose: 0,
		Vol:      0,
		Amount:   0,
		Ztj:      0,
		Zt:       0,
		Dt:       0,
		Dtj:      0,
		Zf:       0,
		Dm:       0,
		Dr:       0,
		Pb:       0,
		A20:      0,
		Lb:       0,
		Prelb:    0,
		St:       0,
		Cy:       0,
		Fcr:      0,
		Fb:       0,
		Ztyz:     0,
		Tp:       0,
		Dtyz:     0,
	}
	t.Log(unsafe.Sizeof(*day))
}
