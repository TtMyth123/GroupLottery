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

type OpenCodeInfo struct {
	//Code       string   `json:"code"`
	//Issue      string   `json:"issue"`
	//Date       string   `json:"date"`
	Jackpots   string   `json:"jackpots"`
	FirstNum   string   `json:"firstNum"`
	SecondNum  []string `json:"secondNum"`
	ThirdNum   []string `json:"thirdNum"`
	ForthNum   []string `json:"forthNum"`
	FifthNum   []string `json:"fifthNum"`
	SixthNum   []string `json:"sixthNum"`
	SeventhNum []string `json:"seventhNum"`
	//EighthNum  []string `json:"eighthNum"`
}

/**
{
	"code": "20200908",
	"issue": "20200908",
	"date": "2020-09-08",
	"week": "TUESDAY",
	"weekZh": "星期二",
	"jackpots": "391022",
	"firstNum": "48355",
	"secondNum": "05747",
	"thirdNum": ["21499", "81691"],
	"forthNum": ["90561", "77660", "48448", "56114", "17580", "82119", "54592"],
	"fifthNum": "2961",
	"sixthNum": ["3348", "6957", "7450"],
	"seventhNum": "421",
	"eighthNum": "78"
}
*/
type OpenCodeInfoNbc struct {
	//Code       string   `json:"code"`
	//Issue      string   `json:"issue"`
	//Date       string   `json:"date"`
	Jackpots   string   `json:"jackpots"`
	FirstNum   string   `json:"firstNum"`
	SecondNum  string   `json:"secondNum"`
	ThirdNum   []string `json:"thirdNum"`
	ForthNum   []string `json:"forthNum"`
	FifthNum   string   `json:"fifthNum"`
	SixthNum   []string `json:"sixthNum"`
	SeventhNum string   `json:"seventhNum"`
	EighthNum  string   `json:"eighthNum"`
}

type OpenCodeZbcInfo struct {
	Nums []string `json:"Nums"`
}
