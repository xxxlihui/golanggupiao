package service

import "math"

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func floatCompare(f1, f2 float32, precision float64) bool {
	return float64Compare(float64(f1), float64(f2), precision)
}
func float64Round(f1, precision float64) float64 {
	return math.Round(f1*math.Pow(10, precision)) * math.Pow(0.1, precision)
}
func floatRound(f1 float32, precision float64) float64 {
	return float64Round(float64(f1), precision)
}
func float64Compare(f1, f2 float64, precision float64) bool {
	n1 := math.Floor(math.Round(f1 * math.Pow(0.1, precision)))
	n2 := math.Floor(math.Round(f2 * math.Pow(0.1, precision)))
	return n1 == n2
}
func zt(preClose, close float32) bool {
	return float64Compare(floatRound(preClose, float64(2)), float64(close), 2)
}
