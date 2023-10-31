package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type responce struct {
	Status  bool
	Message string
	Time    time.Time
	Eroors  interface{}
	Data    interface{}
}

func ErrorMessage(g *gin.Context, code int, message string, errors, data interface{}) {

	res := responce{
		Status:  false,
		Message: message,
		Time:    time.Now(),
		Eroors:  errors,
		Data:    data,
	}

	g.JSON(code, res)
}

func Surcessmessage(g *gin.Context, message string, data interface{}) {
	res := responce{
		Status:  true,
		Message: message,
		Time:    time.Now(),
		Eroors:  nil,
		Data:    data,
	}

	g.JSON(200, res)
}
