package util

import (
	"fmt"
	"testing"
)

func TestMul32(t *testing.T) {
	fmt.Println(FloatMul32(1236, 36, 3))
	fmt.Println(FloatDiv32(1236, 36, 3))
}
func TestFloat32TInt(t *testing.T) {
	fs:=[]float32{12.63,11.23,11.001,11.002,11.123,11.126,11.127}
	for _,v := range fs {
		t.Log(Float32TInt(v,2))
	}
}