package ResultBox

type Xjp28Result struct {
	Data []Xjp28Data `json:"data"`
}
type Xjp28Data struct {
	Opentime  string `json:"opentime"`
	Expect    string `json:"expect"`
	OpenCode  string `json:"opencode"`
	NextTime  string `json:"NextTime"`
	NextIssue string `json:"NextIssue"`
}
