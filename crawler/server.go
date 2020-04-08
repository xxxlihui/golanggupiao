package crawler

import "nn/service"

func GetByCodes(codes []int) []*service.DayRecord {
	return sinaGetByCodes(codes)
}
