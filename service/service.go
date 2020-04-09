package service

import (
	"github.com/gin-gonic/gin"
	"nn/data"
	"nn/log"
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
		rst := GetDB().Raw(
			`select *
from day_records
where day < ?
  and code = ?
order by day desc
limit 1`, r.Day, r.Code).Scan(&preRecord)
		checkDbError(rst.Error)
		r.PreClose = preRecord.Close
		DayAnalyze(r)
		log.Debug("-----保存数据:%+v", r)
		rst = GetDB().Save(&r)
		checkError(rst.Error)
	}
}

func DayStatAnalyze(ctx *gin.Context) {
	/*startday := ctx.GetInt("startDay")
	for{

	}*/

}
