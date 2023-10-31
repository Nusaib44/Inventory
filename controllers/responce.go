package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	Status  bool
	Time    time.Time
	Message string
	Eroors  interface{}
	Data    interface{}
}

func ErrorMessage(g *gin.Context, code int, message string, errors, data interface{}) {

	res := response{
		Status:  false,
		Message: message,
		Time:    time.Now(),
		Eroors:  errors,
		Data:    data,
	}

	g.JSON(code, res)
}

func Surcessmessage(g *gin.Context, message string, data interface{}) {
	res := response{
		Status:  true,
		Message: message,
		Time:    time.Now(),
		Eroors:  nil,
		Data:    data,
	}

	g.JSON(200, res)
}
