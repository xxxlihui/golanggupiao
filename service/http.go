package service

import (
	"github.com/gin-gonic/gin"
)

func StartServer(addr string) error {
	app := gin.Default()
	apiGroup := app.Group("/api")
	apiGroup.POST("/import", ImportData)
	/*apiGroup.GET("/dayStatAnalyze", DayStatAnalyze)
	apiGroup.POST("/getDayStat", GetDayStat)*/
	return app.Run(addr)
}
