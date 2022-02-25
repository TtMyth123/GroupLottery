package Chart

type TypeSchart string

const (
	TypeSchart_Bar TypeSchart = "bar"
)

type TitleSchart struct {
	Text string `json:"text"`
}
type DatasetsSchart struct {
	Label string    `json:"label"`
	Data  []float64 `json:"data"`
}
type OptionsSchart struct {
	Type     TypeSchart       `json:"type"`
	Title    TitleSchart      `json:"title"`
	Labels   []string         `json:"labels"`
	Datasets []DatasetsSchart `json:"datasets"`
}

func NewOptionsSchart(Type TypeSchart, title string, Labels []string, Datasets []DatasetsSchart) OptionsSchart {
	aOptionsSchart := OptionsSchart{
		Type:     TypeSchart_Bar,
		Title:    TitleSchart{Text: title},
		Labels:   Labels,
		Datasets: Datasets,
	}

	return aOptionsSchart
}
