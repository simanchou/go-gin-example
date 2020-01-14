package app

import (
	"github.com/gin-gonic/gin"
	"github.com/simanchou/go-gin-example/pkg/e"
)

//Gin gin struct
type Gin struct {
	C *gin.Context
}

// Response response struct
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}
