package AreaConfig

import (
	"github.com/TtMyth123/GameServer/models"
	"sync"
)

var (
	areaLock      sync.RWMutex
	mpTtRebateSet map[int]RebateSetConfig
)

type RebateSetConfig struct {
	Level       int
	RebateRatio map[int]float64
}

func Init() {
	mpTtRebateSet = make(map[int]RebateSetConfig)

	ReLoadRebateSet(0)
}
func GetRebateSet(area int) RebateSetConfig {
	areaLock.RLock()
	defer areaLock.RUnlock()
	return mpTtRebateSet[area]
}

//func NewRebateSetConfig(area int)RebateSetConfig{
//	aRebateSetConfig := RebateSetConfig{}
//	aRebateSetConfig.RebateRatio = make(map[int]float64)
//	return aRebateSetConfig
//}

func ReLoadRebateSet(area int) {
	areaLock.Lock()
	defer areaLock.Unlock()
	aTtRebateSet := models.GetTtRebateSet()
	aRebateSetConfig := RebateSetConfig{}
	aRebateSetConfig.RebateRatio = make(map[int]float64)
	aRebateSetConfig.Level = aTtRebateSet.Level
	aRebateSetConfig.RebateRatio[1] = aTtRebateSet.BetRebateRatio
	aRebateSetConfig.RebateRatio[2] = aTtRebateSet.BetRebateRatio1
	aRebateSetConfig.RebateRatio[3] = aTtRebateSet.BetRebateRatio2
	aRebateSetConfig.RebateRatio[4] = aTtRebateSet.BetRebateRatio3

	mpTtRebateSet[area] = aRebateSetConfig
}
