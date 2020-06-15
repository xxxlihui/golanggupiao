package dongfangcaifu

import (
	"nn/data"
	"testing"
)

func TestGetReal(t *testing.T) {
	d1 := data.Day{Day: 20200203}
	d2 := data.Day{Day: 20200203}
	t.Logf("d1==d2,%v", d1 == d2)
}
