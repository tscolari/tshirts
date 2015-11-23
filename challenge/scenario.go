package challenge

type Scenario struct {
	ID        string     `json:"scenario_id"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Layers []Layer `json:"layers"`
}

type Layer struct {
	Color  string  `json:"color"`
	Volume float64 `json:"volume"`
}
