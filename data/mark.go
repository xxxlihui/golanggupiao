package data

import "sort"

//市场数据
//日数据
type DayRecord struct {
	PDay
	PCode
	PDayData
	PDaySample
	PDayMood
}

func (receiver *DayRecord) Key() DayCodeKey {
	return DayCodeKey{Day: receiver.Day, Code: receiver.Code}
}

type DayCodeKey struct {
	Day  int
	Code string
}

//时间范围的全部票的日数据
type DayRecords struct {
	DayRecords []*DayRecord
	StartDate  int
	EndDate    int
	//时间索引
	dayMapper map[int][]*DayRecord
	//代码索引 按日期做好排序 倒序
	codeMapper map[string][]*DayRecord
	//时间代码索引
	dayCodeMapper map[DayCodeKey]*DayRecord
}

func (receiver DayRecords) Init(raw []*DayRecord, startTime, endTime int) {
	receiver.DayRecords = raw
	receiver.StartDate = startTime
	receiver.EndDate = endTime
	//按时间排序
	sort.Sort()
}
