package GData

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"sync"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit/ttLog"
)

var aStaffDataManage *StaffDataManage

func InitStaffDataManage() {
	aStaffDataManage = NewStaffDataManage()
}

type StaffData struct {
	Id       int64
	UserName string
	Nickname string
}

type StaffDataManage struct {
	mapStaffData map[int64]*StaffData
	userLock     sync.Mutex
}

func NewStaffDataManage() *StaffDataManage {
	staffDataManage := new(StaffDataManage)
	staffDataManage.mapStaffData = make(map[int64]*StaffData)
	staffDataManage.ReLoadStaffData()
	return staffDataManage
}
func GetStaffDataManage() *StaffDataManage {
	if aStaffDataManage == nil {
		InitStaffDataManage()
	}
	return aStaffDataManage
}

func (this *StaffDataManage) ReLoadStaffData() {
	this.userLock.Lock()
	defer this.userLock.Unlock()

	arrData := make([]StaffData, 0)
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.id, a.user_name, a.nickname from %s a where a.user_type=?`,
		mconst.TableName_TtGameUser)
	c, e := o.Raw(sql, mconst.UserType_4).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
	}
	this.mapStaffData = make(map[int64]*StaffData)
	for _, staff := range arrData {
		this.mapStaffData[staff.Id] = &staff
	}
}
func (this *StaffDataManage) GetStaffData(staffId int64) *StaffData {
	this.userLock.Lock()
	defer this.userLock.Unlock()

	if data, ok := this.mapStaffData[staffId]; ok {
		return data
	} else {
		return nil
	}
}
