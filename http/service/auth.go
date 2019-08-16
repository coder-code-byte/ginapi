package service

import (
	"ginapi/http/handler"
	"ginapi/pkg/exception"
	"ginapi/pkg/logging"
	"ginapi/pkg/setting"
	"ginapi/pkg/utils"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type auth struct {
	LanguageType string  `form:"language_type" binding:"required" json:"language_type"`
	ProductID    int     `form:"product_id" binding:"required" json:"product_id"`
	Time         int     `form:"time" binding:"required" json:"time"`
	TotalAmount  float64 `form:"total_amount" binding:"required" json:"total_amount"`
	TransferID   string  `form:"transfer_id" binding:"required" json:"transfer_id"`
	UserID       string  `form:"user_id" binding:"required" json:"user_id"`
	Sign         string  `form:"sign" binding:"required,len=32" json:"sign"`
}

// Auth this is Auth
func Auth(c *gin.Context) {
	var param auth
	if err := c.ShouldBind(&param); err != nil {
		logging.Error(err.Error())
		handler.SendJSONFail(c, exception.InvalidParams, err.Error())
		c.Abort()
		return
	}
	mp := utils.MapKeyToSnakeCase(structs.Map(param))
	if result := utils.ValidatorSign(mp, setting.APIKEY); result != true {
		handler.SendJSONFail(c, exception.ErrorSign, result)
		c.Abort()
		return
	}
	token, err := utils.GenerateToken(param.LanguageType, param.ProductID, param.UserID, param.TransferID, param.TotalAmount)
	if err != nil {
		logging.Error(err)
		handler.SendJSONFail(c, exception.ErrotAuthToken, err)
		c.Abort()
		return
	}
	var data = make(map[string]string)
	data["token"] = token
	handler.SendJSONSuccess(c, data)
	return
}
