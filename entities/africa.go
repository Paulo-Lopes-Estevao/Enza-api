package entities

type Africa struct {
	Flag        string   `json:"flag"`
	Location    string   `json:"location"`
	Keywords    []string `json:"keywords"`
	Name        string   `json:"name"`
	Capital     string   `json:"capital"`
	Population  float64  `json:"population"`
	Area        float64  `json:"area"`
	Callingcode string   `json:"callingcode"`
}
