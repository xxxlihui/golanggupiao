package service

import (
	"fmt"
	"testing"
)

func TestFloat64Round(t *testing.T) {
	fmt.Printf("%v", Float64Compare(0.0145, 0.0144, 3.0))
}
