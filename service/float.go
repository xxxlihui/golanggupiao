package service

//解决浮点数精度问题
//整数的前三位用作小数位，个位始终保持0做四舍五入用

func pow32(scala int) int32 {
	k := int32(10)
	for i := 0; scala-1 > i; i++ {
		k = k * 10
	}
	return k
}
func pow64(scala int) int64 {
	k := int64(10)
	for i := 0; scala-1 > i; i++ {
		k = k * 10
	}
	return k
}

func Mul32(a, b int32, scala int) int32 {
	return (a*b/pow32(scala-1) + 5) / 10
}
func Div32(a, b int32, scala int) int32 {
	return (a*10*pow32(scala)/b + 5) / 10
}
func Mul64(a, b int64, scala int) int64 {
	return (a*b/(pow64(scala-1)) + 5) / 10
}
func Div64(a, b int64, scala int) int64 {
	return (a*10*pow64(scala)/b + 5) / 10
}
