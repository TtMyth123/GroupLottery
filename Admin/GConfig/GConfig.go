package GConfig

import (
	"encoding/json"
	"io/ioutil"
)

var gConfig *GConfig

type GConfig struct {
	IsDev     bool
	CApiBaseP string //客户端
	SApiBaseP string //服务端
}

type FilmConfig struct {
	MinMoney int
	//Lim
}

func Init() {
	gConfig = new(GConfig)
	confFileName := "conf/GConfig.json"
	gConfig.Reload(confFileName)
}

func (this *GConfig) Reload(fileName string) {
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

func GetGConfig() *GConfig {
	if gConfig == nil {
		Init()
	}
	return gConfig
}
