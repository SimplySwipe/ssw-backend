package models

type ScraperJobInput struct {
	Source       string  `json:"source"`
	ExternalID   string  `json:"externalId"`
	Title        string  `json:"title"`
	Location     string  `json:"location"`
	Description  string  `json:"description"`
	SalaryHourly float64 `json:"salaryHourly"`
	URL          string  `json:"url"`
}
