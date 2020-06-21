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

func NewDefaultDayRecords() *DayRecords {
	return &DayRecords{
		dayMapper:     map[int][]*DayRecord{},
		codeMapper:    map[string][]*DayRecord{},
		dayCodeMapper: map[DayCodeKey]*DayRecord{},
	}
}

func NewDayRecords(raw []*DayRecord, startTime, endTime int) *DayRecords {
	dayRecords := &DayRecords{
		dayMapper:     map[int][]*DayRecord{},
		codeMapper:    map[string][]*DayRecord{},
		dayCodeMapper: map[DayCodeKey]*DayRecord{},
	}
	dayRecords.Init(raw, startTime, endTime)
	return dayRecords
}

//初始化数据和对应的索引
func (receiver *DayRecords) Init(raw []*DayRecord, startTime, endTime int) {
	receiver.DayRecords = raw
	receiver.StartDate = startTime
	receiver.EndDate = endTime
	//按时间排序
	sort.Slice(raw, func(i, j int) bool {
		return raw[i].Day < raw[j].Day
	})
	//建立索引
	for _, record := range raw {
		receiver.dayCodeMapper[record.Key()] = record
		ls := receiver.codeMapper[record.Code]
		if ls == nil {
			ls = []*DayRecord{}
		}
		ls = append(ls, record)
		receiver.codeMapper[record.Code] = ls
		ls = receiver.dayMapper[record.Day]
		if ls == nil {
			ls = []*DayRecord{}
		}
		ls = append(ls, record)
		receiver.dayMapper[record.Day] = ls
	}
}

//根据日期倒序
func (receiver *DayRecords) GetByCode(code string) []*DayRecord {
	ls := receiver.codeMapper[code]
	if ls == nil {
		return []*DayRecord{}
	}
	return ls
}

//获取当天的数据
func (receiver *DayRecords) GetByDay(day int) []*DayRecord {
	ls := receiver.dayMapper[day]
	if ls == nil {
		return []*DayRecord{}
	}
	return ls
}

//获取票的前一个交易日的数据
func (receiver *DayRecords) GetPreDay(code string, day int) *DayRecord {
	days := receiver.GetByCode(code)
	for i, record := range days {
		if record.Day == day {
			if i+1 < len(days) {
				return days[i+1]
			}
		}
	}
	//记录里没有从数据库里加载  todo
	return nil
}
