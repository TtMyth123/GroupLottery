package controllers

import (
	"fmt"
	"github.com/TtMyth123/Admin/controllers/base"
	"github.com/TtMyth123/Admin/controllers/base/enums"
	"github.com/TtMyth123/kit"
	"github.com/TtMyth123/kit/httpKit"
	"os"
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
