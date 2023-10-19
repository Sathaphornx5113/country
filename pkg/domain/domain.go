package domain

type CountryData struct {
	Country    string      `json:"country"`
	Year       int32       `json:"year"`
	Region     interface{} `json:"region"`
	Population int64       `json:"population"`
}
