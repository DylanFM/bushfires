package main

import (
	"time"
)

// Returns a slice of IncidentFeatures. The slice contains the latest report for all incidents marked as current
func currentIncidentsWithLatestReport() (incidents []IncidentFeature, err error) {
	// Select the latest report for all current incidents
	stmt, err := db.Prepare(`SELECT DISTINCT ON (i.uuid) r.uuid, incident_uuid,
                              guid, title, link, category, pubdate,
                              updated, alert_level, location, council_area, status, fire_type,
                              fire, size, responsible_agency, extra,
                              ST_Y(r.geometry) as lat, ST_X(r.geometry) as lng
                            FROM incidents i
                            JOIN reports r ON i.uuid = r.incident_uuid
                            WHERE i.current = true
                            ORDER BY i.uuid, r.created_at DESC`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}

	for rows.Next() {
		var lat string
		var lng string
		var ip IncidentProperties
		var fea IncidentFeature

		err = rows.Scan(&ip.ReportUUID, &fea.UUID, &ip.Guid, &ip.Title, &ip.Link, &ip.Category, &ip.Pubdate, &ip.Updated, &ip.AlertLevel, &ip.Location, &ip.CouncilArea, &ip.Status, &ip.FireType, &ip.Fire, &ip.Size, &ip.ResponsibleAgency, &ip.Extra, &lat, &lng)
		if err != nil {
			return
		}

		// Report properties needs to be placed within a IncidentFeature
		// and the IncidentFeature needs to be adjusted based on ip content
		fea.Type = "Feature"
		fea.Properties = ip
		fea.Geometry = pointFromCoordinates(lng, lat)

		incidents = append(incidents, fea)
	}

	return
}

// Takes a slice of report features and returns it wrapped in a GeoJSON FeatureCollection
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
                              guid, title, link, category, pubdate,
                              updated, alert_level, location, council_area, status, fire_type,
                              fire, size, responsible_agency, extra,
                              ST_Y(r.geometry) as lat, ST_X(r.geometry) as lng
                            FROM incidents i
                            JOIN reports r ON i.uuid = r.incident_uuid
                            WHERE i.current = true
                            AND i.uuid = $1
                            ORDER BY i.uuid, r.created_at DESC`)
	if err != nil {
		return fea, err
	}
	defer stmt.Close()

	var lat string
	var lng string
	var ip IncidentProperties

	err = stmt.QueryRow(uuid).Scan(&ip.ReportUUID, &ip.Guid, &ip.Title, &ip.Link, &ip.Category, &ip.Pubdate, &ip.Updated, &ip.AlertLevel, &ip.Location, &ip.CouncilArea, &ip.Status, &ip.FireType, &ip.Fire, &ip.Size, &ip.ResponsibleAgency, &ip.Extra, &lat, &lng)
	if err != nil {
		return fea, err
	}

	// Incident properties needs to be placed within a ReportFeature
	// and the IncidentFeature needs to be adjusted based on ip content
	fea.Type = "Feature"
	fea.Properties = ip
	fea.UUID = uuid
	fea.Geometry = pointFromCoordinates(lng, lat)

	return fea, nil
}

// Takes a lat and lng and returns a GeoJSON Point with those coordinates
func pointFromCoordinates(lat string, lng string) (p Point) {
	p = Point{}

	p.Type = "Point"
	p.Coordinates = []string{lat, lng}

	return
}

// IncidentFeature is a GeoJSON feature for an individual report
type IncidentFeature struct {
	Type       string             `json:"type"`
	UUID       string             `json:"id"`
	Geometry   Point              `json:"geometry"`
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
	Updated           time.Time `json:"updated"`
	AlertLevel        string    `json:"alertLevel"`
	Location          string    `json:"location"`
	CouncilArea       string    `json:"councilArea"`
	Status            string    `json:"status"`
	FireType          string    `json:"fireType"`
	Fire              bool      `json:"fire"`
	Size              string    `json:"size"`
	ResponsibleAgency string    `json:"responsibleAgency"`
	Extra             string    `json:"extra"`
	//Polygons          []string
}

// Point is a lng, lat coordinate pair used for the GeoJSON geometry type Point.
type Point struct {
	Type        string   `json:"type"`
	Coordinates []string `json:"coordinates"`
}
