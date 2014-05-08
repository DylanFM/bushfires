HOST: http://api.bushfir.es/

# Fire incidents API
An API to current and previous [RFS](http://www.rfs.nsw.gov.au) major incidents.

# GET /incidents
Retrieve a collection of incidents which may or may not be filtered by parameters included in the request.

+ Parameters

    + current = `true` (optional, boolean) ... Boolean value indicating whether to only include current incidents or not.

+ Response 200 (application/json)

  JSON conforming to the GeoJSON spec. Response is a feature collection containing 0 or more features. Each feature represents an incident. Properties included in the feature object are taken from the incident's most recent report.

  + Body

            {
              "type": "FeatureCollection",
              "features": [
                {
                  "type": "Feature",
                  "id": "690e3280-6d61-480c-b7db-83dbaf923b25",
                  "geometry": {
                    "type": "Point",
                    "coordinates": [
                      "153.614",
                      "-28.7052"
                    ]
                  },
                  "properties": {
                    "reportUuid": "049e5351-15e9-410a-86bf-ea4ca6b315f5",
                    "guid": "tag:www.rfs.nsw.gov.au,2014-04-05:157939",
                    "title": "Broken Head Reserve Rd, Broken Head",
                    "link": "http:\/\/www.rfs.nsw.gov.au\/dsp_content.cfm?cat_id=683",
                    "category": "Not Applicable",
                    "pubdate": "0001-01-01T10:04:52+10:04",
                    "updated": "2014-04-05T22:31:00+11:00",
                    "alertLevel": "Not Applicable",
                    "location": "184 Broken Head Reserve Rd, Broken Head, NSW 2481",
                    "councilArea": "Byron",
                    "status": "being controlled",
                    "fireType": "Backyard\/Bbq\/HungiB\/onfire\/Cooking fire",
                    "fire": true,
                    "size": "0 ha",
                    "responsibleAgency": "Rural Fire Service",
                    "extra": ""
                  }
                },
                {
                  "type": "Feature",
                  "id": "df34a81e-4fe5-4c09-8403-466526ea503c",
                  "geometry": {
                    "type": "Point",
                    "coordinates": [
                      "152.4392",
                      "-32.3934"
                    ]
                  },
                  "properties": {
                    "reportUuid": "1b44ceb8-e3b9-4a5d-8965-af6992efe892",
                    "guid": "tag:www.rfs.nsw.gov.au,2014-03-05:155600",
                    "title": "The Lakes Way, Bungwahl",
                    "link": "http:\/\/www.rfs.nsw.gov.au\/dsp_content.cfm?cat_id=683",
                    "category": "Advice",
                    "pubdate": "0001-01-01T10:04:52+10:04",
                    "updated": "2014-03-05T20:04:00+11:00",
                    "alertLevel": "Advice",
                    "location": "1826 The Lakes Way, Bungwahl, NSW 2423",
                    "councilArea": "Great Lakes",
                    "status": "under control",
                    "fireType": "Grass fire",
                    "fire": true,
                    "size": "4 ha",
                    "responsibleAgency": "Rural Fire Service",
                    "extra": ""
                  }
                }
              ]
            }

# GET /incidents/{id}
Retrieve an incident by its *id*.

+ Parameters

    + id (required, string, `df34a81e-4fe5-4c09-8403-466526ea503c`) ... Id of an incident.

+ Response 200 (application/json)

  JSON conforming to the GeoJSON spec. Response is a feature collection containing 1 feature. The feature represents the requested incident. Properties included in the feature object are taken from the incident's most recent report.

  + Body

            {
              "type": "FeatureCollection",
              "features": [
                {
                  "type": "Feature",
                  "id": "df34a81e-4fe5-4c09-8403-466526ea503c",
                  "geometry": {
                    "type": "Point",
                    "coordinates": [
                      "152.4392",
                      "-32.3934"
                    ]
                  },
                  "properties": {
                    "reportUuid": "1b44ceb8-e3b9-4a5d-8965-af6992efe892",
                    "guid": "tag:www.rfs.nsw.gov.au,2014-03-05:155600",
                    "title": "The Lakes Way, Bungwahl",
                    "link": "http:\/\/www.rfs.nsw.gov.au\/dsp_content.cfm?cat_id=683",
                    "category": "Advice",
                    "pubdate": "0001-01-01T10:04:52+10:04",
                    "updated": "2014-03-05T20:04:00+11:00",
                    "alertLevel": "Advice",
                    "location": "1826 The Lakes Way, Bungwahl, NSW 2423",
                    "councilArea": "Great Lakes",
                    "status": "under control",
                    "fireType": "Grass fire",
                    "fire": true,
                    "size": "4 ha",
                    "responsibleAgency": "Rural Fire Service",
                    "extra": ""
                  }
                }
              ]
            }


# GET /incidents/{id}/reports
Retrieve all of an incident's reports, identifying the incident by its *id*.

+ Parameters

    + id (required, string, `df34a81e-4fe5-4c09-8403-466526ea503c`) ... Id of an incident.

+ Response 200 (application/json)

  JSON conforming to the GeoJSON spec. Response is a feature collection containing 1 or more features. Each feature represents a report for the requested incident. The features are ordered by *updated* date ascending.

  + Body

            { ... }

