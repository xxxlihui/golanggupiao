package service

import (
	"nn/data"
	"nn/log"
	"time"
)

var dayRecordsCache map[string][]*data.DayRecord

func LoadDaysToCache() {
	//先清理缓存
	dayRecordsCache = make(map[string][]*data.DayRecord, 3800)
	lc, err := time.LoadLocation("Asia/Shanghai")
	checkError(err)
	now := time.Now().Add(100 * time.Hour * 24 * -1).In(lc)
	rs := make([]*data.DayRecord, 0, 3800*100)
	GetDB().Order("day").Where("day>=?", now.Year()*10000+int(now.Month())*100+now.Day()).Find(&rs)
	for _, r := range rs {
		addDayRecords(r)
	}
	log.Info("加载完缓存")
}
func addDayRecords(r *data.DayRecord) {
	if rs, ex := dayRecordsCache[r.Code]; ex {
		rs = append(rs, r)
		dayRecordsCache[r.Code] = rs
	} else {
		rs := make([]*data.DayRecord, 0, 100)
		rs = append(rs, r)
		dayRecordsCache[r.Code] = rs
	}
}
