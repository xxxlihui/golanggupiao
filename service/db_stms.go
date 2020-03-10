package service
const (
	_ StmType = iota
	existsDayStmt
	getDayStmt
	saveDayStmt
)

var stms = map[StmType]*Stm{
	existsDayStmt: &Stm{Sql: "select 1 from day_record where id=$1"},
	getDayStmt:    &Stm{Sql: "select id,day,code,high,low,colse,amount,vol,zt,dt,dm,dr,pb,stop,lb from day_record where day=$1 and code=$2 where id=$1"},
	saveDayStmt: &Stm{Sql: `insert into day_record
(id, day, code, high, low, close, amount, vol, zt, dt, dm, dr, pb, stop, lb)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`},
}
