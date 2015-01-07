package main

import (
	"encoding/json"
	"time"
)

// MinimalIncidentFeatureCollection A GeoJSON FeatureCollection representing a collection of incidents with minimal details
type MinimalIncidentFeatureCollection struct {
	Type     string                   `json:"type"`
	Features []MinimalIncidentFeature `json:"features"`
}

type MinimalIncidentFeature struct {
	Type       string                    `json:"type"`
	UUID       string                    `json:"id"`
	Geometry   json.RawMessage           `json:"geometry"`
	Properties MinimalIncidentProperties `json:"properties"`
}

// MinimalIncidentProperties is a limited collection of properties for an individual report included in a MinimalIncidentFeature GeoJSON response.
type MinimalIncidentProperties struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	FireType string `json:"fireType"`
}

// Takes a slice of minimal incident features and returns it wrapped in a GeoJSON FeatureCollection
func minimalIncidentFeatureCollectionForMinimalIncidentFeatures(fea []MinimalIncidentFeature) MinimalIncidentFeatureCollection {
	return MinimalIncidentFeatureCollection{"FeatureCollection", fea}
}

// Returns a slice of IncidentFeatures. The incidents are those within the specified time range
func minimalIncidentsWithinTimeRange(timeStart time.Time, timeEnd time.Time) (incidents []MinimalIncidentFeature, err error) {

	// ST_CollectionExtract will filter the collection and return its points
	stmt, err := db.Prepare(`SELECT DISTINCT ON (i.uuid) incident_uuid,
                              title, link, fire_type,
                              ST_AsGeoJSON(ST_CollectionExtract(geometry, 1))
                            FROM incidents i
                            JOIN reports r ON i.uuid = r.incident_uuid
                            WHERE current_from && tstzrange($1, $2)
                            ORDER BY i.uuid, r.pubdate DESC`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(timeStart.UTC().Format(time.RFC3339), timeEnd.UTC().Format(time.RFC3339))
	if err != nil {
		return
	}

	for rows.Next() {
		var ip MinimalIncidentProperties
		var fea MinimalIncidentFeature
		var geom string

		err = rows.Scan(&fea.UUID, &ip.Title, &ip.Link, &ip.FireType, &geom)
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
