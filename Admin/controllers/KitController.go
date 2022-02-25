package controllers

import (
	"fmt"
	"os"
	"ttmyth123/GroupLottery/Admin/controllers/base"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
	"ttmyth123/kit"
	"ttmyth123/kit/httpKit"
)

type KitController struct {
	base.ABaseController
}

func (c *KitController) GetGuid() {
	a := kit.GetGuid()
	c.JsonResult(enums.JRCodeSucc, "", a)
}

func (c *KitController) UploadPic() {
	picGuid := c.GetString("picGuid")
	imgFile, head, err := c.GetFile("file")
	filePath := ""
	if err == nil {
		fileName := picGuid
		filePath, _ = httpKit.UploadTmpFile(60*60, imgFile, head, "static/upload/tmp", fileName, "")
		fileInfo, _ := os.Stat(filePath)
		a := fileInfo.Name()
		fmt.Println(a)

	}
	c.JsonResult(enums.JRCodeSucc, "", filePath)
}
