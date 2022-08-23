package controllers

import (
	"github.com/TtMyth123/Admin/controllers/base/enums"
	"github.com/TtMyth123/kit"
)

func (c *ApiController) GetGuid() {
	a := kit.GetGuid()
	c.JsonResult(enums.JRCodeSucc, "", a)
}
