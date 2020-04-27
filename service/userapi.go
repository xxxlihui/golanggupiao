package service

import (
	"github.com/gin-gonic/gin"
	"nn/data"
)

func GetDayStat(context *gin.Context) {
	param := &struct {
		StartTime int64 `json:"startTime"`
	}{}
	checkError(context.BindJSON(&param))
	dayStats := make([]data.DayStat, 0)
	GetDB().Where("day>=?", param.StartTime).Find(&dayStats)
	context.JSON(200, dayStats)
}
