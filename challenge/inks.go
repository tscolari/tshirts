package challenge

type InksResponse struct {
	Inks Inks `json:"inks"`
}

type Inks []Ink

type Ink struct {
	ID    string  `json:"id"`
	Color string  `json:"color"`
	Cost  float64 `json:"cost"`
}
