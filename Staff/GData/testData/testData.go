package testData

import "ttmyth123/kit/timeKit"

var (
	ChatId      int64
	arrGameData []GameData
)

type GameData struct {
	GameId   int
	RoomId   int
	UserName string
	ChatId   int64
}

func init() {
	ChatId = 100
	arrGameData = make([]GameData, 4)
	arrGameData[0] = GameData{GameId: 11, RoomId: 100, UserName: "张三"}
	arrGameData[1] = GameData{GameId: 12, RoomId: 101, UserName: "李四"}
	arrGameData[2] = GameData{GameId: 13, RoomId: 102, UserName: "王五"}
	arrGameData[3] = GameData{GameId: 14, RoomId: 103, UserName: "赵六"}
}
func GetChatId() int64 {
	ChatId++
	return ChatId
}

func GetGameData() GameData {
	i := timeKit.GlobaRand.Int31n(3)
	arrGameData[i].ChatId = GetChatId()
	return arrGameData[i]
}
