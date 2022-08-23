package ArticleBll

import (
	"fmt"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"time"
)

type Notice struct {
	Id          int
	Seq         int
	Title       string
	TitleImgUrl string
	Des         string
	State       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetArticleList(ArticleType, state, pageIndex, pageSize int) (int, []Notice) {
	o := orm.NewOrm()
	PageTotal := 0
	arrData := make([]Notice, 0)
	sqlArgs := make([]interface{}, 0)
	sqlWhere := `where a.article_type=?`
	sqlArgs = append(sqlArgs, ArticleType)
	if state != 0 {
		sqlWhere = sqlWhere + ` and a.state=?`
		sqlArgs = append(sqlArgs, state)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from %s a %s`,
		mconst.TableName_TtArticleInfo, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
	if e != nil {
		return 0, arrData
	}

	offset, _ := sqlKit.GetOffset(PageTotal, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.seq desc LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.* from %s a %s`,
		mconst.TableName_TtArticleInfo, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		return 0, arrData
	}

	return PageTotal, arrData
}

type ArticleInfoEx struct {
	Id int
	//Title       string
	//HtmlContent string

	Title           string `orm:"size(512)"`  //标题
	TitleImgUrl     string `orm:"size(512)"`  //标题图片
	Des             string `orm:"size(5120)"` //概要（描述）
	HtmlContentFile string `orm:"size(512)"`  //内容
	ArticleType     int
}

func GetArticleInfo(id int) (ArticleInfoEx, error) {
	aArticleInfoEx := ArticleInfoEx{}
	articleInfo, e := models.GetArticleInfo(id)
	if e != nil {
		return aArticleInfoEx, e
	}
	aArticleInfoEx.Id = id
	aArticleInfoEx.Title = articleInfo.Title
	aArticleInfoEx.TitleImgUrl = articleInfo.TitleImgUrl
	aArticleInfoEx.Des = articleInfo.Des

	if articleInfo.HtmlContentFile != "" {
		fileT, e := ioutil.ReadFile(articleInfo.HtmlContentFile)
		if e != nil {
			return aArticleInfoEx, e
		}
		aArticleInfoEx.HtmlContentFile = string(fileT)
	}

	return aArticleInfoEx, nil
}

func DelArticleInfo(id int) error {
	o := orm.NewOrm()
	aYhArticleInfo := models.TtArticleInfo{Id: id}
	_, e := o.Delete(&aYhArticleInfo)
	return e
}
