package main

import (
	"encoding/json"
	"time"
)

// Returns a slice of IncidentFeatures. The slice contains the latest report for all incidents marked as current
func currentIncidentsWithLatestReport() (incidents []IncidentFeature, err error) {
	// Select the latest report for all current incidents
	stmt, err := db.Prepare(`SELECT DISTINCT ON (i.uuid) r.uuid, incident_uuid,
                              guid, title, link, category, timezone('UTC', pubdate),
                              alert_level, location, council_area, status, fire_type,
                              fire, size, responsible_agency, extra,
                              timezone('UTC', lower(i.current_from)) as first_seen,
                              timezone('UTC', upper(i.current_from)) as last_seen,
                              ST_AsGeoJSON(geometry)
                            FROM incidents i
                            JOIN reports r ON i.uuid = r.incident_uuid
                            WHERE i.current = true
                            ORDER BY i.uuid, r.pubdate DESC`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}

	for rows.Next() {
		var ip IncidentProperties
		var fea IncidentFeature
		var geom string

		err = rows.Scan(&ip.ReportUUID, &fea.UUID, &ip.Guid, &ip.Title, &ip.Link, &ip.Category, &ip.Pubdate, &ip.AlertLevel, &ip.Location, &ip.CouncilArea, &ip.Status, &ip.FireType, &ip.Fire, &ip.Size, &ip.ResponsibleAgency, &ip.Extra, &ip.FirstSeen, &ip.LastSeen, &geom)
		if err != nil {
			return
		}

		// Report properties needs to be placed within a IncidentFeature
		// and the IncidentFeature needs to be adjusted based on ip content
		fea.Type = "Feature"
		fea.Properties = ip
		fea.Geometry = json.RawMessage([]byte(geom))

		incidents = append(incidents, fea)
	}

	return
}

// Takes a slice of incident features and returns it wrapped in a GeoJSON FeatureCollection
func incidentFeatureCollectionForIncidentFeatures(fea []IncidentFeature) IncidentFeatureCollection {
	return IncidentFeatureCollection{"FeatureCollection", fea}
}

// IncidentFeatureCollection A GeoJSON FeatureCollection representing a collection of ReportFeatures
type IncidentFeatureCollection struct {
	Type     string            `json:"type"`
	Features []IncidentFeature `json:"features"`
}

// Takes a UUID and returns a IncidentFeature representing that incident
func incidentFeatureForUUID(uuid string) (IncidentFeature, error) {
	fea := IncidentFeature{}

	// Select the latest report for the requested incident
	stmt, err := db.Prepare(`SELECT DISTINCT ON (i.uuid) r.uuid,
                              guid, title, link, category, timezone('UTC', pubdate),
                              alert_level, location, council_area, status, fire_type,
                              fire, size, responsible_agency, extra,
                              timezone('UTC', lower(i.current_from)) as first_seen,
                              timezone('UTC', upper(i.current_from)) as last_seen,
                              ST_AsGeoJSON(geometry)
                            FROM incidents i
                            JOIN reports r ON i.uuid = r.incident_uuid
                            WHERE i.current = true
                            AND i.uuid = $1
                            ORDER BY i.uuid, r.created_at DESC`)
	if err != nil {
		return fea, err
	}
	defer stmt.Close()

	var ip IncidentProperties
	var geom string

	err = stmt.QueryRow(uuid).Scan(&ip.ReportUUID, &ip.Guid, &ip.Title, &ip.Link, &ip.Category, &ip.Pubdate, &ip.AlertLevel, &ip.Location, &ip.CouncilArea, &ip.Status, &ip.FireType, &ip.Fire, &ip.Size, &ip.ResponsibleAgency, &ip.Extra, &ip.FirstSeen, &ip.LastSeen, &geom)
	if err != nil {
		return fea, err
	}

	// Incident properties needs to be placed within a ReportFeature
	// and the IncidentFeature needs to be adjusted based on ip content
	fea.Type = "Feature"
	fea.Properties = ip
	fea.UUID = uuid
	fea.Geometry = json.RawMessage([]byte(geom))

	return fea, nil
}

// IncidentFeature is a GeoJSON feature for an individual incident (but mostly filled with data from its latest report)
type IncidentFeature struct {
	Type       string             `json:"type"`
	UUID       string             `json:"id"`
	Geometry   json.RawMessage    `json:"geometry"`
	Properties IncidentProperties `json:"properties"`
}

// IncidentProperties is a collection of properties for an individual report included in a IncidentFeature GeoJSON response.
type IncidentProperties struct {
	ReportUUID        string    `json:"reportUuid"`
	Guid              string    `json:"guid"`
	Title             string    `json:"title"`
	Link              string    `json:"link"`
	Category          string    `json:"category"`
	Pubdate           time.Time `json:"pubdate"`
	FirstSeen         time.Time `json:"firstSeen"`
	LastSeen          time.Time `json:"lastSeen"`
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

// A GeoJSON FeatureCollection representing a collection of ReportFeatures
type ReportFeatureCollection struct {
	Type     string          `json:"type"`
	Features []ReportFeature `json:"features"`
}
