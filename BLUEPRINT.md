HOST: http://api.bushfir.es/

# Fire incidents API
An API to current and previous [RFS](http://www.rfs.nsw.gov.au) major incidents.

# GET /incidents
Retrieve a collection of incidents which may or may not be filtered by parameters included in the request.

+ Parameters

    + current = `true` (optional, boolean) ... Boolean value indicating whether to only include current incidents or not.

+ Response 200 (application/json)

  JSON conforming to the GeoJSON spec. Response is a feature collection containing 0 or more features. Each feature represents an incident. Properties included in the feature object are taken from the incident's most recent report.

        {"type":"FeatureCollection","features":[{"type":"Feature","id":"690e3280-6d61-480c-b7db-83dbaf923b25","geometry":{"type":"Point","coordinates":["153.614","-28.7052"]},"properties":{"reportUuid":"049e5351-15e9-410a-86bf-ea4ca6b315f5","guid":"tag:www.rfs.nsw.gov.au,2014-04-05:157939","title":"Broken Head Reserve Rd, Broken Head","link":"http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=683","category":"Not Applicable","pubdate":"0001-01-01T10:04:52+10:04","updated":"2014-04-05T22:31:00+11:00","alertLevel":"Not Applicable","location":"184 Broken Head Reserve Rd, Broken Head, NSW 2481","councilArea":"Byron","status":"being controlled","fireType":"Backyard/Bbq/HungiB/onfire/Cooking fire","fire":true,"size":"0 ha","responsibleAgency":"Rural Fire Service","extra":""}},{"type":"Feature","id":"ecbc1a72-3883-4fea-aa76-579d3695a1f8","geometry":{"type":"Point","coordinates":["149.3854","-36.9118"]},"properties":{"incidentUuid":"9f2f40ff-2fdb-4589-82eb-4b693d9a697c","guid":"tag:www.rfs.nsw.gov.au,2014-04-05:157724","title":"Jones Boundary Windrow Burn","link":"http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=683","category":"Advice","pubdate":"0001-01-01T10:04:52+10:04","updated":"2014-04-03T10:45:00+11:00","alertLevel":"Advice","location":"Intersection of Jones Boundary Rd and Coolangubra Forest Way ","councilArea":"Bombala","status":"under control","fireType":"Hazard Reduction","fire":true,"size":"70 ha","responsibleAgency":"Forests NSW","extra":""}},{"type":"Feature","id":"b5beaff8-41f6-49c0-8607-8c8d2c71f451","geometry":{"type":"Point","coordinates":["149.1464","-37.1466"]},"properties":{"incidentUuid":"acf7632d-411b-4519-ba01-0454bc22e851","guid":"tag:www.rfs.nsw.gov.au,2014-04-05:157285","title":"Bondi Camp Windrow Burn","link":"http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=683","category":"Advice","pubdate":"0001-01-01T10:04:52+10:04","updated":"2014-04-02T09:20:00+11:00","alertLevel":"Advice","location":"22 km south of Bombala, Bondi State Forest","councilArea":"Bombala","status":"under control","fireType":"Hazard Reduction","fire":true,"size":"225 ha","responsibleAgency":"Forests NSW","extra":""}},{"type":"Feature","id":"e521dd35-bdfd-4dfd-a9be-58b548026c98","geometry":{"type":"Point","coordinates":["149.706","-33.9888"]},"properties":{"incidentUuid":"ae0e9f29-c069-460d-87d5-b5a273ed2bc8","guid":"tag:www.rfs.nsw.gov.au,2014-04-05:157555","title":"Arkstone Plantation Establishment Burn","link":"http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=683","category":"Advice","pubdate":"0001-01-01T10:04:52+10:04","updated":"2014-04-02T08:30:00+11:00","alertLevel":"Advice","location":"Vulcan State Forest: 35kms south-west of Oberon in vicinity of Arkstone Rd, The Blue Rd and Tooles Rd. ","councilArea":"Oberon","status":"under control","fireType":"Hazard Reduction","fire":true,"size":"370 ha","responsibleAgency":"Forests NSW","extra":""}},{"type":"Feature","id":"c3d22d41-b225-4deb-b124-5d87e29b6cce","geometry":{"type":"Point","coordinates":["151.6034","-33.1567"]},"properties":{"incidentUuid":"bfb7b432-4d54-4202-b49a-5542300b5822","guid":"tag:www.rfs.nsw.gov.au,2014-04-05:84431","title":"Crangan Bay-West Coal Fire","link":"http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=683","category":"Advice","pubdate":"0001-01-01T10:04:52+10:04","updated":"2014-04-05T10:00:00+11:00","alertLevel":"Advice","location":"500 metres west of Pacific Hwy at Crangan Bay. Same location as previous Incident - No 13113083624.","councilArea":"Lake Macquarie","status":"under control","fireType":"Scrub fire","fire":true,"size":"0 ha","responsibleAgency":"NSW National Parks and Wildlife Service","extra":""}},{"type":"Feature","id":"e49afe3d-727c-4a9b-8bc1-74a09a848094","geometry":{"type":"Point","coordinates":["150.5386","-29.5656"]},"properties":{"incidentUuid":"caae8556-35a6-4e96-8d2a-b7ba56c8bdf5","guid":"tag:www.rfs.nsw.gov.au,2014-04-05:157915","title":"Bingara Rd, Warialda","link":"http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=683","category":"Not Applicable","pubdate":"0001-01-01T10:04:52+10:04","updated":"2014-04-05T22:53:00+11:00","alertLevel":"Not Applicable","location":"Bingara Rd, Warialda, NSW 2402","councilArea":"Gwydir","status":"under control","fireType":"Structure fire (A fire involving a residential, commercial or industrial building)","fire":true,"size":"0 ha","responsibleAgency":"Rural Fire Service","extra":""}}]}

# GET /incidents/{id}
Retrieve an incident by its *id*.

+ Parameters

    + id (required, string, `df34a81e-4fe5-4c09-8403-466526ea503c`) ... Id of an incident.

+ Response 200 (application/json)

  JSON conforming to the GeoJSON spec. Response is a feature collection containing 1 feature. The feature represents the requested incident. Properties included in the feature object are taken from the incident's most recent report.

  + Body

        {"type":"FeatureCollection","features":[{"type":"Feature","id":"df34a81e-4fe5-4c09-8403-466526ea503c","geometry":{"type":"Point","coordinates":["152.4392","-32.3934"]},"properties":{"reportUuid":"1b44ceb8-e3b9-4a5d-8965-af6992efe892","guid":"tag:www.rfs.nsw.gov.au,2014-03-05:155600","title":"The Lakes Way, Bungwahl","link":"http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=683","category":"Advice","pubdate":"0001-01-01T10:04:52+10:04","updated":"2014-03-05T20:04:00+11:00","alertLevel":"Advice","location":"1826 The Lakes Way, Bungwahl, NSW 2423","councilArea":"Great Lakes","status":"under control","fireType":"Grass fire","fire":true,"size":"4 ha","responsibleAgency":"Rural Fire Service","extra":""}}]}


# GET /incidents/{id}/reports
Retrieve all of an incident's reports, identifying the incident by its *id*.

+ Parameters

    + id (required, string, `df34a81e-4fe5-4c09-8403-466526ea503c`) ... Id of an incident.

+ Response 200 (application/json)

  JSON conforming to the GeoJSON spec. Response is a feature collection containing 1 or more features. Each feature represents a report for the requested incident. The features are ordered by *updated* date ascending.

        { ... }

