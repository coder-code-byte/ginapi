package handler

import (
	"ginapi/pkg/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response this is Response
type Response struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendJSONSuccess this is SendJSONSuccess
func SendJSONSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  true,
		Code:    200,
		Message: exception.GetMsg(200),
		Data:    data,
	})
}

// SendJSONFail this is SendJSONFail
func SendJSONFail(c *gin.Context, code int, errMsg interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  false,
		Code:    code,
		Message: exception.GetMsg(code),
		Data:    errMsg,
	})
}
