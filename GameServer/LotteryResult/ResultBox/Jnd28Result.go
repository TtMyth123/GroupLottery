package ResultBox

type Game28Result struct {
	NextIssue string
	NextTime  string
	Expect    string   `json:"expect"`
	Opencode  []string `json:"opencode"`
	Opentime  string   `json:"opentime"`
	R         int
}
