package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nn/data"
	"nn/log"
	"sort"
)

func ImportData(context *gin.Context) {
	log.Debug("请求importData")
	records := make([]*data.DayRecord, 0)
	if err := context.BindJSON(&records); err != nil {
		checkError(err)
	}
	for _, r := range records {
		log.Info("导入:%+v", r)
		//取上一条记录
		preRecord := &data.DayRecord{}
		rst := data.GetDB().Raw(
			`select *
from day_records
where day < ?
  and code = ?
order by day desc
limit 1`, r.Day, r.Code).Scan(&preRecord)
		checkDbError(rst.Error)
		r.PreClose = preRecord.Close
		r.Prelb = preRecord.Lb
		DayAnalyze(r)
		log.Debug("-----保存数据:%+v", r)
		if r.Day == 20200410 && r.Code == "sh600030" {
			fmt.Printf("---")
		}
		rst = data.GetDB().Save(&r)
		checkError(rst.Error)
	}
}

//触发分析数据，以后这个不用处理
func DayStatAnalyze(ctx *gin.Context) {
	startday := ctx.GetInt("startDay")
	analyzeDay(startday)
}

func analyzeDay(startDay int) {
	//统计非连板数据
	sql1 := `insert
into day_stats
(day, zt, dt, dm, pb, a20, fb, dr, fcr, tp, ztyz, dtyz,z,d)
select day,
       sum(zt)   zt,
       sum(dt)   dt,
       sum(dm)   dm,
       sum(pb)   pb,
       sum(a20)  a20,
       sum(fb)   fb,
       sum(dr)   dr,
       sum(fcr)  fcr,
       sum(tp)   tp,
       sum(ztyz) ztyz,
       sum(dtyz) dtyz,
       sum(case when zf>0 then 1 else 0 end) z,
       sum(case when zf<0 then 1 else 0 end) d
from day_records
where day >= ?
group by day
order by day
on conflict (day)
    do update set zt=excluded.zt,
                  pb=excluded.pb,
                  dt=excluded.dt,
                  a20=excluded.a20,
                  fcr=excluded.fcr,
                  fb=excluded.fb,
                  ztyz=excluded.ztyz,
                  tp=excluded.tp,
                  dtyz=excluded.dtyz,
                  dm=excluded.dm,
                  dr=excluded.dr,
                  z=excluded.z,
                  d=excluded.d`
	rs := data.GetDB().Exec(sql1, startDay)
	checkDbError(rs.Error)
	//统计连板数据,加入市场最高板和市场第二高板
	sql2 := `select sum(1) c,day,lb
from day_records
where day>=? and lb>0
group by day,lb
order by day ,lb
`
	type d struct {
		C   int
		Day int
		Lb  int
	}
	ds := make([]*d, 0)
	rst := data.GetDB().Raw(sql2, startDay).Find(&ds)
	checkDbError(rst.Error)
	gds := make(map[int][]*d)
	//group 分组
	for _, v := range ds {
		dd := gds[v.Day]
		if dd == nil {
			dd = make([]*d, 0)
		}
		dd = append(dd, v)
		gds[v.Day] = dd
	}
	for k, v := range gds {
		//排序
		sort.Slice(v, func(i, j int) bool {
			return v[i].Day < v[j].Day
		})
		dst := &data.DayStat{}
		rst = data.GetDB().Model(&dst).Where("day=?", k).Scan(&ds)
		checkError(rst.Error)
		if len(v) == 0 {
			updateMax(0, 0, 0, 0, k)
		} else if len(v) == 1 {
			vv := v[0]
			updateMax(vv.Lb, vv.C, vv.Lb, vv.C, k)
		} else {
			vvMax := v[len(v)-1]
			vvMax2 := v[len(v)-2]
			updateMax(vvMax.Lb, vvMax.C, vvMax2.Lb, vvMax2.C, k)
		}
	}

}
func updateMax(mx, mxn, mx2, mx2n, day int) {
	rst := data.GetDB().Exec("update day_stats set mx=?,mxn=?,mx2=?,mx2n=? where day=?",
		mx, mxn, mx2, mx2n, day,
	)
	checkError(rst.Error)
}
