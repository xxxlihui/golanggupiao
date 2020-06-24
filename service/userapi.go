package service

import (
	"github.com/gin-gonic/gin"
	"nn/data"
)

func GetDayStat(context *gin.Context) {
	param := &struct {
		StartTime int64 `json:"start"`
	}{}
	checkError(context.BindJSON(&param))
	dayStats := make([]data.DayStat, 0)
	data.GetDB().Order("day").Where("day>=?", param.StartTime).Find(&dayStats)
	context.JSON(200, dayStats)
}
