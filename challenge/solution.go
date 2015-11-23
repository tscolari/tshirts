package challenge

type Solution struct {
	ScenarioID string   `json:"scenario_id"`
	Answers    []Answer `json:"answers"`
}

type Answer struct {
	Inks []string `json:"inks"`
}
