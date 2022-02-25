package ResultBox

//type Xjp28Result struct {
//	Data []Xjp28Data `json:"data"`
//}
//type Xjp28Data struct {
//	Opentime  string `json:"opentime"`
//	Expect    string `json:"expect"`
//	OpenCode  string `json:"opencode"`
//	NextTime  string `json:"NextTime"`
//	NextIssue string `json:"NextIssue"`
//}

/**
{
    "Data": {
        "Arr": [
            {
                "Issue": 218842,
                "Result": "6,7,1",
                "Time": "2021-01-21 12:30:00"
            }
        ],
        "NextIssue": "218862",
        "NextTime": "2021-01-21 13:30:00",
        "Timer": 112
    },
    "Msg": "",
    "Result": true
}
*/

type XG28Result struct {
	Data   XG28Data
	Msg    string
	Result bool
}
type XG28Code struct {
	Issue  string
	Result string
	Time   string
}

type XG28Data struct {
	Arr       []XG28Code
	NextIssue string
	NextTime  string
	Timer     int
}
