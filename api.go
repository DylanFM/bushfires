package main

func reportFeatureCollectionForReportFeatures(rf []ReportFeature) (rfc ReportFeatureCollection) {
	rfc = ReportFeatureCollection{"FeatureCollection", rf}

	return
}

// A GeoJSON FeatureCollection representing a collection of ReportFeatures
type ReportFeatureCollection struct {
	Type     string          `json:"type"`
	Features []ReportFeature `json:"features"`
}

// Takes a UUID and returns a ReportFeature representing that report
func reportFeatureForUUID(uuid string) (rf ReportFeature, err error) {
	rf = ReportFeature{}

	stmt, err := db.Prepare(`SELECT title, ST_Y(geometry) as lat, ST_X(geometry) as lng FROM reports WHERE uuid = $1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	var lat string
	var lng string
	var rp ReportProperties
	err = stmt.QueryRow(uuid).Scan(&rp.Title, &lat, &lng)
	if err != nil {
		// err very well may be sql.ErrNoRows which says that no rows matched the uuid
		return
	}

	// Report properties needs to be placed within a ReportFeature
	// and the ReportFeature needs to be adjusted based on rp content
	rf.Type = "Feature"
	rf.Properties = rp
	rf.UUID = uuid
	rf.Geometry = pointFromCoordinates(lng, lat)

	return
}

// Takes a lat and lng and returns a GeoJSON Point with those coordinates
func pointFromCoordinates(lat string, lng string) (p Point) {
	p = Point{}

	p.Type = "Point"
	p.Coordinates = []string{lat, lng}

	return
}

// ReportFeature is a GeoJSON feature for an individual report
type ReportFeature struct {
	Type       string           `json:"type"`
	UUID       string           `json:"id"`
	Geometry   Point            `json:"geometry"`
	Properties ReportProperties `json:"properties"`
}

// ReportProperties is a collection of properties for an individual report included in a ReportFeature GeoJSON response.
type ReportProperties struct {
	// IncidentUUID      string
	// Hash              string
	// Guid              string
	Title string `json:"title"`
	//Link              string
	//Category          string
	//Pubdate           time.Time
	//Description       string
	//Updated           time.Time
	//AlertLevel        string
	//Location          string
	//CouncilArea       string
	//Status            string
	//FireType          string
	//Fire              bool
	//Size              string
	//ResponsibleAgency string
	//Extra             string
	//Points            string // Geojson
	//Polygons          []string
	//CreatedAt         time.Time
	//UpdatedAt         time.Time
}

// Point is a lng, lat coordinate pair used for the GeoJSON geometry type Point.
type Point struct {
	Type        string   `json:"type"`
	Coordinates []string `json:"coordinates"`
}
