package service

import (
	"github.com/gin-gonic/gin"
)

func addDay(context *gin.Context) {
	day := &DayRecord{}
	if err := context.BindJSON(&day); err != nil {
		panic("解析数据错误")
	}
	rst := GetDB().Where("day=? and code=?", day.Day, day.Code)
	checkError(rst.Error)
	if rst.RowsAffected > 0 {
		panic("记录已经存在")
	}
	//计算 计算字段
	//拿上一条记录来
	preDay := &DayRecord{}
	rst = GetDB().Where("code=?", day.Code).Order("day desc").First(&preDay)
	checkError(rst.Error)
	if rst.RowsAffected == 0 {
		//没有上次的记录
		//直接插入数据
		GetDB().Save(&day)
		return
	}

}
