package data

type Data struct {
	Records   map[string]*OneRecordDays //全部的日期数据
	StartDate int                       //开始日期
	CurDate   int                       //当前日期
}

//根据代码返回记录
func (this *Data) getRecordsByCode(code string) *OneRecordDays {
	return this.Records[code]
}
func NewData(startDate, curDate int, records []*DayRecord) *Data {
	d := make(map[string][]*DayRecord)
	for _, v := range records {
		code := v.Code
		rs := d[code]
		if rs == nil {
			rs = make([]*DayRecord, 0)
		}
		rs = append(rs, v)
		d[code] = rs
	}
	data := &Data{StartDate: startDate, CurDate: curDate, Records: make(map[string]*OneRecordDays)}
	for code, v := range d {
		data.Records[code] = NewOneRecordDays(v)
	}
	return data
}
