package TtHint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ttmyth123/kit/stringKit"
)

type CN struct {
	C string
	N string
}
type TtHint struct {
	dict map[string]CN
}

func NewTtHint(path, fileName string) *TtHint {
	gTtHint := new(TtHint)
	confFileName := fmt.Sprintf("%s/%s.json", path, fileName)
	gTtHint.Reload(confFileName)
	return gTtHint
}

func (this *TtHint) Reload(fileName string) {
	//读取用户自定义配置
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &this.dict)
	if err != nil {
		panic(err)
	}
}

func (this *TtHint) GetHint(hint string) string {
	if str, ok := this.dict[hint]; ok {
		return str.C
	} else {
		return hint
	}
}

func (this *TtHint) GetMpString(mp ...interface{}) string {
	return stringKit.GetJsonStr(mp)
}
