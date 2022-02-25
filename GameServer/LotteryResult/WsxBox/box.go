package WsxBox

/**
{
    "row": 1,
    "code": "bbc",
    "data": [
        {
            "opentime": "2020-08-05 17:15:00",
            "expect": "20200805",
            "opencode": "{\"code\":\"20200805\",\"issue\":\"20200805\",\"date\":\"2020-08-05\",\"week\":\"WEDNESDAY\",\"weekZh\":\"星期三\",\"jackpots\":\"714690\",\"firstNum\":\"05072\",\"secondNum\":\"56151\",\"thirdNum\":[\"55974\",\"82260\"],\"forthNum\":[\"87186\",\"28791\",\"12550\",\"56521\",\"12168\",\"25889\",\"39503\"],\"fifthNum\":\"1072\",\"sixthNum\":[\"0182\",\"2525\",\"3308\"],\"seventhNum\":\"368\",\"eighthNum\":\"06\"}",
            "NextTime": "2020-08-06 17:15:00",
            "NextIssue": "20200806"
        }
    ],
    "SleepTime": 0,
    "LastDataStr": "",
    "LastIssue": "",
    "SetTimerIng": false,
    "LastMsg": {},
    "Lock": {}
}
*/

/**
{
    "code": "zbc",
    "data": [
        {
            "opentime": "2021-01-14 13:10:00",
            "expect": "202101141310",
            "opencode": "{\"2\":\"14\",\"5\":\"29\",\"16\":\"71\",\"18\":\"79\",\"19\":\"80\",\"0\":\"02\",\"7\":\"42\",\"11\":\"51\",\"12\":\"56\",\"14\":\"63\",\"1\":\"09\",\"4\":\"20\",\"6\":\"31\",\"9\":\"44\",\"17\":\"78\",\"3\":\"16\",\"8\":\"43\",\"10\":\"48\",\"13\":\"57\",\"15\":\"65\"}",
            "Timer": 336,
            "NextTime": "2021-01-15 13:20:00",
            "NextIssue": "202101151320",
            "NextTime2": "2021-01-15 13:30:00",
            "NextIssue2": "202101151330"
        }
    ],
    "ServerTime": "2021-01-14 13:14:21"
}
*/

type WsxResultData struct {
	Opentime   string `json:"opentime"`
	Expect     string `json:"expect"`
	Opencode   string `json:"opencode"`
	NextTime   string `json:"NextTime"`
	NextIssue  string `json:"NextIssue"`
	NextTime2  string `json:"NextTime2"`
	NextIssue2 string `json:"NextIssue2"`
	ServerTime string `json:"ServerTime"`
}

type WsxResultInfo struct {
	Row        int    `json:"row"`
	Code       string `json:"code"`
	Data       []WsxResultData
	ServerTime string
}
