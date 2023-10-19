package domain

type CountryData struct {
	Country    string `json:"country"`
	Year       int32  `json:"year"`
	Region     string `json:"region"`
	Population int32  `json:"population"`
}
