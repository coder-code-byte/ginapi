package v1

import (
	"ginapi/http/handler"
	"ginapi/pkg/exception"
	"ginapi/pkg/logging"

	"github.com/gin-gonic/gin"
)

type test struct {
	Username string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

// Test this is Test
func Test(c *gin.Context) {
	var te test
	if err := c.ShouldBind(&te); err != nil {
		logging.Error(err.Error())
		handler.SendJSONFail(c, exception.InvalidParams, err.Error())
		return
	}
	handler.SendJSONSuccess(c, te)
	return
}

type player struct {
	Username string `form:"username" json:"username" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
}

// TestPost this is TestPost
func TestPost(c *gin.Context) {
	var pla player
	if err := c.ShouldBind(&pla); err != nil {
		logging.Error(err.Error())
		handler.SendJSONFail(c, exception.InvalidParams, err.Error())
		return
	}
	handler.SendJSONSuccess(c, pla)
	return
}
