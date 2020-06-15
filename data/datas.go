package data

type Data struct {
	Records   map[string]*OneRecordDays //全部的日期数据
	Infos     map[string]*RecordInfo    //上市信息
	StartDate int                       //开始日期
	CurDate   int                       //当前日期
}

//根据代码返回记录
func (this *Data) GetRecordsByCode(code string) *OneRecordDays {
	return this.Records[code]
}
func (this *Data) GetInfoByCode(code string) *RecordInfo {
	return this.Infos[code]
}
func (this *Data) LoadData(startDate, curDate int, records []*DayRecord, infos []*RecordInfo) {
	data := this
	data.StartDate = startDate
	data.CurDate = curDate
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
	for code, v := range d {
		data.Records[code] = NewOneRecordDays(v)
	}
	for _, info := range infos {
		data.Infos[info.Code] = info
	}
}

//获取上市日期
func (this *Data) GetLaunchDate(code string) (int, error) {
	info := this.Infos[code]
	if info == nil {
		return 0, NoInfoError
	}
	return info.LaunchDate, nil
}

//获取上市天数
/*func (this *Data) GetDayCount(code string) (int, error) {
	launchData, err := this.GetLaunchDate(code)
	if err != nil {
		return 0, err
	}

}*/
func EmptyData() *Data {
	return &Data{
		Records:   make(map[string]*OneRecordDays),
		Infos:     make(map[string]*RecordInfo),
		StartDate: 0,
		CurDate:   0,
	}
}
func NewData(startDate, curDate int, records []*DayRecord, infos []*RecordInfo) *Data {
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
	for _, info := range infos {
		data.Infos[info.Code] = info
	}
	return data
}
