package util

import "time"

//整数类型的日期转换成时间类型
func FromIntDay(day int) time.Time {
	return time.Date(day/10000, time.Month(day%10000/100), day%100,
		0, 0, 0, 0,
		time.Local,
	)
}

//计算两个整数日期之间相差的天数 day1-day2
func DiffIntDay(day1, day2 int) int {
	t1 := FromIntDay(day1)
	t2 := FromIntDay(day2)
	return int(t1.Sub(t2).Hours() / 24)
}

//计算now-day相差的天数
func DiffNowIntDay(day int) int {
	return int(time.Now().Sub(FromIntDay(day)).Hours() / 24)
}
