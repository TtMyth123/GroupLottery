账号，昵称，推荐人，推荐id，上分，下分，洗码，盈利，上分下分洗码盈利的开始统计时间，是否推手，是否测试账号，头像地址，余额。


http://47.244.125.83:7710/GetResult/WSX1     南部彩
http://47.244.125.83:7710/GetResult/WSX2     北部彩
http://47.244.125.83:7710/GetResult/WSX3     中部彩

{
	"code": "20200805",
	"issue": "20200805",
	"date": "2020-08-05",
	"week": "WEDNESDAY",
	"weekZh": "星期三",
	"jackpots": "714690",
	"firstNum": "05072",
	"secondNum": "56151",
	"thirdNum": ["55974", "82260"],
	"forthNum": ["87186", "28791", "12550", "56521", "12168", "25889", "39503"],
	"fifthNum": "1072",
	"sixthNum": ["0182", "2525", "3308"],
	"seventhNum": "368",
	"eighthNum": "06"
}
#规则
- 头特 

0~9:jackpots:714690 倒数第2位数字。如:9

- 尾特

0~9:jackpots 最后1位数字。如:0

- 头等特码

00~99 jackpots  最后两位

- 一等特码

00~99 firstNum  最后两位

- 二等特码
 
00~99 secondNum  最后两位

- 二连位

00~99 全部号码 最后两位

- 三连位

000~999 全部号码 最后三位

- 平码两位区
00~99,买两个号码如：53,23  所有开奖号码最后两位，出现视为中奖

- 平码三位区

000~999 ,买三个号码如：01,53,23  所有开奖号码最后两位，出现视为中奖

- 波色
00~99 jackpots  最后两位

大小单双 特码
jackpots 最后两个数

大小单双 一等特 
firstNum 最后两个数






//管理员列表
GET /Chat/LF_WK/GetAdminList
//建立与某个用户的私聊
GET /Chat/LF_WK/CreatePrivateTo
//拿历史消息
GET /Chat/LF_WK/GetMsgData
//拿房间用户信息
GET /Chat/LF_WK/GetRoomInfo
//后台聊天
POST /Chat/LF_WK/AdminMsg
//建立管理员账号
POST /Chat/LF_WK/CreateAdminAccount
