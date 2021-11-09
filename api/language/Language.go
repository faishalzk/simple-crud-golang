package language

type Language struct {
	Language string `json:"language"`
	Appeared int64 `json:"appeared"`
	Functional bool	`json:"functional"`
	ObjectOriented bool `json:"object-oriented"`
	Created []string `json:"created"`
	Relation *Relation `json:"relation"`
}

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences []string `json:"influences"`
}
