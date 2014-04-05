package main

// CountResponse includes the number of incidents and number of reports present in the database.
type CountResponse struct {
	IncidentCount int `json:"incidentCount"`
	ReportCount   int `json:"reportCount"`
}
