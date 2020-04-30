package service

import (
	"fmt"
	"testing"
)

func TestMul32(t *testing.T) {
	fmt.Println(FloatMul32(1236, 36, 3))
	fmt.Println(FloatDiv32(1236, 36, 3))
}
