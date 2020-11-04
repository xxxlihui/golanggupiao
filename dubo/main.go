package main

import "fmt"

func main() {
	threadCount := 4 //线程数
	total := 10000   //总投资金额
	//把total分成线程数等份
	max := total / threadCount
	less := total % threadCount
	xfr := 1.42 //特朗普赔率
	yfr := 3.25 //拜登赔率
	xf := int(xfr * 100)
	yf := int(yfr * 100)
	for c := 0; c < threadCount; c++ {
		start := c * max
		end := start + max
		if c == threadCount-1 {
			end = end + less
		}
		start++
		fmt.Printf("start:%d,end：%d---%d,%d", start, end, xf, yf)
	}
}

func enums(start, end, total, xf, yf int) [][4]int {
	rst := make([][4]int, 0, (end-start)/2)
	for x := start; x <= end; x++ {
		for y := start; y <= end && (x+y) <= total; y++ {
			xz := x*xf - y*100
			yz := y*yf - x*100
			if xz > 0 && yz > 0 {
				rst = append(rst, [4]int{x, y, xz, yz})
			}
		}
	}
	return rst
}
