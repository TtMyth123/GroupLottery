package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/GConfig"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/astaxie/beego/orm"
	"time"
)

type TtArticleInfo struct {
	Id              int
	Seq             int    //序号
	Title           string `orm:"size(512)"`  //标题
	TitleImgUrl     string `orm:"size(512)"`  //标题图片
	Des             string `orm:"size(5120)"` //概要（描述）
	HtmlContentFile string `orm:"size(512)"`  //内容
	ArticleType     int
	ArticleTypeDes  string `orm:"size(200)"`

	State     int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (d TtArticleInfo) MarshalJSON() ([]byte, error) {
	type Alias TtArticleInfo

	StrTime1 := ""
	StrTime2 := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrTime1 = d.CreatedAt.Format(timeKit.DateTimeLayout)
		StrTime2 = d.UpdatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrTime1 = d.CreatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
		StrTime2 = d.UpdatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		CreatedAt string
		UpdatedAt string
	}{
		Alias:     (Alias)(d),
		CreatedAt: StrTime1,
		UpdatedAt: StrTime2,
	})
}

type NoticeInfo struct {
	Id          int
	Title       string    //标题
	TitleImgUrl string    //标题图片
	Des         string    //概要（描述）
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

func (d NoticeInfo) MarshalJSON() ([]byte, error) {
	type Alias NoticeInfo

	StrTime := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrTime = d.UpdatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrTime = d.UpdatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		StrUpdatedAt string
	}{
		Alias:        (Alias)(d),
		StrUpdatedAt: StrTime,
	})
}

func (this *TtArticleInfo) TableName() string {
	return mconst.TableName_TtArticleInfo
}

func InitArticleInfoList() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_TtArticleInfo).Count()
	if c == 0 {
		arr := make([]TtArticleInfo, 0)

		arr = append(arr, TtArticleInfo{Id: 1000, State: mconst.ArticleState_2_Enabled, ArticleType: mconst.ArticleType_1_GG, ArticleTypeDes: mconst.GetArticleTypeName(mconst.ArticleType_1_GG), Title: mconst.GetArticleTypeName(mconst.ArticleType_1_GG)})

		o.InsertMulti(len(arr), arr)
	}
	return nil
}

func (this *TtArticleInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.State = mconst.ArticleState_2_Enabled
	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = int(id)
	return e
}

func (this *TtArticleInfo) Update(o orm.Ormer, cols ...string) error {
	this.UpdatedAt = time.Now()
	_, e := o.Update(this, cols...)
	return e
}
func GetArticleInfo(id int) (TtArticleInfo, error) {
	o := orm.NewOrm()
	aTtArticleInfo := TtArticleInfo{Id: id}
	e := o.Read(&aTtArticleInfo)
	if e != nil {
		return aTtArticleInfo, errors.New("没有对应的数据")
	}
	return aTtArticleInfo, nil
}

func GetNoticeList(lastId, c int) ([]NoticeInfo, int) {
	o := orm.NewOrm()
	arrData := make([]NoticeInfo, 0)
	sqlArgs := make([]interface{}, 0)

	if c == 0 {
		c = 10
	}

	sqlWhere := "where a.state=? and a.article_type=?"
	sqlArgs = append(sqlArgs, mconst.ArticleState_2_Enabled, mconst.ArticleType_1_GG)

	if lastId != 0 {
		sqlWhere += " and a.id < ?"
		sqlArgs = append(sqlArgs, lastId)
	}

	sql := fmt.Sprintf(`select a.* from
yh_article_info a %s order by a.seq desc, a.id desc limit ?`,
		sqlWhere)
	sqlArgs = append(sqlArgs, c)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)
	newLastId := 0
	iLen := len(arrData)
	if iLen > 0 {
		newLastId = arrData[iLen-1].Id
	}

	return arrData, newLastId
}
