package crawler

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGetByCodes(t *testing.T) {
	var u big.Float

	fmt.Printf("value:%s\n", u.String())
	fmt.Printf("%+v\n", u)

}

type U struct {
	Name string
	F    int
}
