package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}, msg interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
		"msg":  msg,
	})

}

func Fail(c *gin.Context, data interface{}, msg interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 1,
		"data": data,
		"msg":  msg,
	})

}
func Response(c *gin.Context, code int, data interface{}, msg interface{}) {
	c.JSON(code, gin.H{
		"code": 0,
		"data": data,
		"msg":  msg,
	})
}
