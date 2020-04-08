package service

import (
	"github.com/gin-gonic/gin"
)

func StartHttp(port string) {
	app := gin.Default()

	app.Run(port)
}

type Err struct {
	Msg string `json:"msg"`
}
