package service

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"testing"
)

func TestFloat64Round(t *testing.T) {
	d := decimal.Decimal{}
	t.Logf("%+v", d)
	str := `{
  "n1": 23.023,
  "n2": 11
}`
	n := N{}

	json.Unmarshal([]byte(str), &n)
	t.Logf("%+v", n)
}

type N struct {
	N1 decimal.Decimal `json:"n1"`
	N2 decimal.Decimal `json:"n2"`
}
