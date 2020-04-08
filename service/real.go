package service

import (
	"fmt"
	"nn/crawler"
	"nn/data"
	"runtime/debug"
	"strings"
)

var follows = make([]*data.DayRecord, 0)
var all = make(map[string]*data.DayRecord)
var allCode = make([]string, 0)
var curIndex = 0
var lenPerOne = 200

func newDayRecord(record *data.DayRecord) {
	cur := all[record.Code]
	cur.Close = record.Close
	cur.Amount = record.Amount
	cur.Vol = record.Vol
	cur.High = record.High
	cur.Low = record.Low
	cur.Open = record.Open
	cur.Name = record.Name
	//分析数据
	DayAnalyze(cur)
	//保存数据到数据库
	GetDB().Save(cur)
}

var startGetData = false

func StartGetData() {
	if startGetData {
		return
	}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("获取数据失败:%v", p)
				debug.PrintStack()
			}
		}()
		first := curIndex
		last := curIndex + 200
		if first > len(allCode) {
			first = 0
			last = curIndex + 200
		}
		if last > len(allCode) {
			last = len(allCode)
		}
		codes := allCode[first:last]
		records, err := crawler.GetByCodes(codes)
		if err != nil {
			fmt.Printf("获取codes:%s失败\n", strings.Join(codes, ","))
		}
		for _, e := range records {
			newDayRecord(e)
		}
	}()
}

var startAnalyzeDayStat = false

func StartAnalyzeDayStat() {
	if startAnalyzeDayStat {
		return
	}
	go func() {

	}()

}
