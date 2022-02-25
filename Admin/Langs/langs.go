package Langs

import (
	"encoding/json"
	"io/ioutil"
)

var gLangs *Langs

type Langs struct {
	Lang  string
	Langs map[string]map[int]string
}

func Init() {
	gLangs = new(Langs)
	confFileName := "conf/langs.json"
	gLangs.Reload(confFileName)
}

func (this *Langs) Reload(fileName string) {
	//读取用户自定义配置
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		panic(err)
	}
}

func GetLangs() *Langs {
	if gLangs == nil {
		Init()
	}
	return gLangs
}
