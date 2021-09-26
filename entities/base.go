package entities

type Base struct {
	Flag        string   `json:"flag"`
	Location    string   `json:"location"`
	Keywords    []string `json:"keywords"`
	Name        string   `json:"name"`
	Capital     string   `json:"capital"`
	Currency    string   `json:"currency"`
	Language    string   `json:"language"`
	Population  float64  `json:"population"`
	Area        float64  `json:"area"`
	Callingcode string   `json:"callingcode"`
}
