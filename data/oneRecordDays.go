package data

import "sort"

type OneRecordDays struct {
	DayRecords []*DayRecord
}

//获取当日的值
func (this *OneRecordDays) getByDay(day int) *DayRecord {
	for _, r := range this.DayRecords {
		if r.Day == day {
			return r
		}
	}
	return nil
}

//前一日的值
func (this *OneRecordDays) getPreByDay(day int) *DayRecord {
	for i, r := range this.DayRecords {
		if r.Day == day {
			if len(this.DayRecords)-1 < i+1 {
				return nil
			} else {
				return this.DayRecords[i+1]
			}
		}
	}
	return nil
}

//设置当日的值
func (this *OneRecordDays) setByDay(day int, record *DayRecord) {
	for i, r := range this.DayRecords {
		if r.Day == day {
			this.DayRecords[i] = record
		}
	}
}
func NewOneRecordDays(records []*DayRecord) *OneRecordDays {
	sort.Slice(records, func(i, j int) bool {
		return records[i].Day > records[j].Day
	})
	return &OneRecordDays{DayRecords: records}
}
