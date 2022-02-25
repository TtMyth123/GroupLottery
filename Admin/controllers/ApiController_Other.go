package controllers

import (
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
	"ttmyth123/kit"
)

func (c *ApiController) GetGuid() {
	a := kit.GetGuid()
	c.JsonResult(enums.JRCodeSucc, "", a)
}
