package WsxBox

/**
{
	"code": "20200908",
	"issue": "20200908",
	"date": "2020-09-08",
	"week": "TUESDAY",
	"weekZh": "星期二",
	"jackpots": "03827",
	"firstNum": "58985",
	"secondNum": ["61297", "74667"],
	"thirdNum": ["05374", "94559", "81972", "35892", "96178", "16374"],
	"forthNum": ["8009", "1502", "4937", "5002"],
	"fifthNum": ["2645", "4547", "0886", "4310", "6104", "0097"],
	"sixthNum": ["018", "167", "324"],
	"seventhNum": ["50", "73", "21", "98"]
}
*/
type HttpOpenCodeBbc struct {
	Jackpots   string   `json:"0"`
	FirstNum   string   `json:"1"`
	SecondNum  []string `json:"2"`
	ThirdNum   []string `json:"3"`
	ForthNum   []string `json:"4"`
	FifthNum   []string `json:"5"`
	SixthNum   []string `json:"6"`
	SeventhNum []string `json:"7"`
	//EighthNum  []string `json:"8"`
}
