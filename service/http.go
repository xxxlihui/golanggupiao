package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartHttp(port string) {
	app := gin.Default()
	app.GET("/api/addDay", func(context *gin.Context) {
		dayRecord := &DayRecord{}
		if err := context.BindJSON(dayRecord); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, &Err{Msg: "数据格式错误"})
			return
		}

	})

	app.Run(port)
}

type Err struct {
	Msg string `json:"msg"`
}
