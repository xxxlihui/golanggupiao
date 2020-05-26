package util

import (
	"testing"
)

func TestFromIntDay(t *testing.T) {
	t1 := FromIntDay(20200501)
	t.Log(t1)
}

func TestDiffIntDay(t *testing.T) {
	d := DiffIntDay(20200809, 20200102)
	t.Log(d)
}

func TestDiffNowIntDay(t *testing.T) {
	d := DiffNowIntDay(20200501)
	t.Log(d)
}
