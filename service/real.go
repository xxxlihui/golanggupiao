package service

var follows = make([]*DayRecord, 0)
var all = make(map[int]*DayRecord)
var allCode = make([]int, 0)
var curIndex = 0
var lenPerOne = 20

func newDayRecord(record *DayRecord) {
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
