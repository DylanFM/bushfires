package main

import (
	"encoding/json"
	"time"
)

type ReportFeatureCollection struct {
	Type     string          `json:"type"`
	Features []ReportFeature `json:"features"`
}

// A GeoJSON FeatureCollection representing a collection of ReportFeatures
// ReportFeature is a GeoJSON feature for an individual report
type ReportFeature struct {
	Type       string           `json:"type"`
	UUID       string           `json:"id"`
	Geometry   json.RawMessage  `json:"geometry"`
	Properties ReportProperties `json:"properties"`
}

// ReportProperties is a collection of properties for an individual report included in a ReportFeature GeoJSON response.
type ReportProperties struct {
	Guid              string    `json:"guid"`
	Title             string    `json:"title"`
	Link              string    `json:"link"`
	Category          string    `json:"category"`
	Pubdate           time.Time `json:"pubdate"`
	AlertLevel        string    `json:"alertLevel"`
	Location          string    `json:"location"`
	CouncilArea       string    `json:"councilArea"`
	Status            string    `json:"status"`
	FireType          string    `json:"fireType"`
	Fire              bool      `json:"fire"`
	Size              string    `json:"size"`
	ResponsibleAgency string    `json:"responsibleAgency"`
	Extra             string    `json:"extra"`
}

// Takes a slice of report features and returns it wrapped in a GeoJSON FeatureCollection
func reportFeatureCollectionForReportFeatures(fea []ReportFeature) ReportFeatureCollection {
	return ReportFeatureCollection{"FeatureCollection", fea}
}

// Takes an incident's UUID and returns the incident's reports as a slice of ReportFeature
func reportsForIncident(uuid string) (reports []ReportFeature, err error) {
	// Select all reports for the incident, ordered from oldest to newest
	stmt, err := db.Prepare(`SELECT uuid,
                              guid, title, link, category, timezone('UTC', pubdate),
                              alert_level, location, council_area, status, fire_type,
                              fire, size, responsible_agency, extra,
                              ST_AsGeoJSON(geometry)
                            FROM reports
                            WHERE incident_uuid = $1
                            ORDER BY pubdate ASC`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(uuid)
	if err != nil {
		return
	}

	for rows.Next() {
		var rp ReportProperties
		var fea ReportFeature
		var geom string

		err = rows.Scan(&fea.UUID, &rp.Guid, &rp.Title, &rp.Link, &rp.Category, &rp.Pubdate, &rp.AlertLevel, &rp.Location, &rp.CouncilArea, &rp.Status, &rp.FireType, &rp.Fire, &rp.Size, &rp.ResponsibleAgency, &rp.Extra, &geom)
		if err != nil {
			return
		}

		// Report properties needs to be placed within a ReportFeature
		// and the ReportFeature needs to be adjusted based on rp content
		fea.Type = "Feature"
		fea.Properties = rp
		fea.Geometry = json.RawMessage([]byte(geom))

		reports = append(reports, fea)
	}

	return
}
