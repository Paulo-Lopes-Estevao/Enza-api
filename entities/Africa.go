package main

type Country struct {
	Flag        string  `json:"flag"`
	Location    string  `json:"location"`
	Name        string  `json:"name"`
	Capital     string  `json:"capital"`
	Population  float64 `json:"population"`
	Area        float64 `json:"area"`
	Callingcode int32   `json:"callingcode"`
}
