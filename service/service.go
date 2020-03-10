package service

import "database/sql"

func addDay(record *DayRecord) {

}
func saveDay(record *DayRecord) {

}
func getDay(day int, code int) (error, *DayRecord) {
	dayRecord := &DayRecord{}
	stm := stms[saveDayStmt].Stm
	stm.QueryRow(getId(day, code)).Scan(&dayRecord.Id,
		&dayRecord.Day, &dayRecord.Code, &dayRecord.High,
		&dayRecord.Low, &dayRecord.Close, &dayRecord.Amount,
		&dayRecord.Vol, &dayRecord.Zt, &dayRecord.Dt,
		&dayRecord.Dm, &dayRecord.Dr, && dayRecord.Pb,
		&dayRecord.Stop, &dayRecord.Lb)
	return nil, nil
}

type StmType int
type Stm struct {
	Sql string
	Stm *sql.Stmt
}

func getId(day, code int) int64 {
	return int64(day*1000000 + code)
}

func existsDay(day int, code int) (error, bool) {
	stm := stms[existsDayStmt].Stm
	c := 0
	if err := stm.QueryRow(getId(day, code)).Scan(&c); err != nil {
		return err, false
	}
	return nil, c > 0
}

func initStms() {
	for _, v := range stms {
		if stm, err := db.Prepare(v.Sql); err != nil {
			panic(`sql:select 1 from day_record where day=$1 and code=$2 错误` + err.Error())
		} else {
			v.Stm = stm
		}
	}

}
func closeStms() {
	for _, v := range stms {
		if v.Stm != nil {
			v.Stm.Close()
			v.Stm = nil
		}
	}
}
