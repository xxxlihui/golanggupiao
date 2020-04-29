package service

import (
	"fmt"
	"testing"
)

func TestMul32(t *testing.T) {
	fmt.Println(Mul32(1236, 36, 3))
	fmt.Println(Div32(1236, 36, 3))
}
