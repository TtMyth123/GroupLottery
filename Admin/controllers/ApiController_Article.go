package controllers

import (
	"github.com/astaxie/beego/orm"
	"os"
	//"ttmyth123/DummyMarket/YhFilmAdmin/controllers/cBll"
	"ttmyth123/GroupLottery/Admin/controllers/ArticleBll"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/kit/httpKit/TmpFileKit"
)

func (c *ApiController) GetArticleList() {
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 10)
	state, _ := c.GetInt("state", 0)
	articleType, _ := c.GetInt("articleType", 0)
	PageTotal, arrData := ArticleBll.GetArticleList(articleType, state, pageIndex, pageSize)
	c.JsonPageResult(enums.JRCodeSucc, "", PageTotal, arrData)
}

func (c *ApiController) DelArticle() {
	id, _ := c.GetInt("Id")
	e := ArticleBll.DelArticleInfo(id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, "", "")
	}

	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) GetArticle() {
	id, _ := c.GetInt("Id")
	aArticleInfo, e := ArticleBll.GetArticleInfo(id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, "", "")
	}

	c.JsonResult(enums.JRCodeSucc, "", aArticleInfo)
}

func (c *ApiController) SaveNotice() {
	id, _ := c.GetInt("Id")

	aArticleInfo, e := models.GetArticleInfo(id)
	if id != 0 {
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "")
		}
	}
	aArticleInfo.Title = c.GetString("Title")
	aArticleInfo.Des = c.GetString("Des")
	aArticleInfo.State, _ = c.GetInt("State", 0)
	aArticleInfo.Seq, _ = c.GetInt("Seq", 0)

	if c.GetString("TitleImgUrl") == "" {
		os.Remove(aArticleInfo.TitleImgUrl)
		aArticleInfo.TitleImgUrl = ""
	}
	o := orm.NewOrm()
	if aArticleInfo.Id == 0 {
		aArticleInfo.ArticleType, _ = c.GetInt("ArticleType", mconst.ArticleType_1_GG)
		aArticleInfo.State, _ = c.GetInt("State", mconst.ArticleState_2_Enabled)
		e = aArticleInfo.Add(o)
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "")
		}
	} else {
		e = aArticleInfo.Update(o, "Seq", "Title", "State", "Des", "UpdatedAt")
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "")
		}
	}

	c.JsonResult(enums.JRCodeSucc, "", aArticleInfo)
}

func (c *ApiController) SaveArticleHasTmpPic() {
	id, _ := c.GetInt("Id")
	PicGuid := c.GetString("PicGuid")

	aArticleInfo, e := models.GetArticleInfo(id)
	if id != 0 {
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "")
		}
	}
	aArticleInfo.Title = c.GetString("Title")
	aArticleInfo.Des = c.GetString("Des")
	aArticleInfo.State, _ = c.GetInt("State", 0)
	aArticleInfo.Seq, _ = c.GetInt("Seq", 0)

	if PicGuid != "" {
		filePath, e := TmpFileKit.ToNewPath(PicGuid, `static/upload/Article`, "")
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "没有图片数据")
		}

		aArticleInfo.TitleImgUrl = filePath
	}
	o := orm.NewOrm()
	if aArticleInfo.Id == 0 {
		aArticleInfo.ArticleType, _ = c.GetInt("ArticleType", mconst.ArticleType_1_GG)
		aArticleInfo.State, _ = c.GetInt("State", mconst.ArticleState_2_Enabled)
		e = aArticleInfo.Add(o)
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "")
		}
	} else {
		e = aArticleInfo.Update(o, "Seq", "Title", "State", "Des", "UpdatedAt", "TitleImgUrl")
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "")
		}
	}

	c.JsonResult(enums.JRCodeSucc, "", aArticleInfo)
}
