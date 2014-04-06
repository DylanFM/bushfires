# Bushfires

_Part of a collection of tools used to provide an API to NSW bushfire data_: [Data collector](https://github.com/dylanfm/major-incidents-data), [Importer](https://github.com/DylanFM/incident-worker) and [GeoJSON API (this repo)](https://github.com/DylanFM/bushfires)

The NSW Rural Fire Service publish [RSS feeds](http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=1358) of incident data. In particular, they provide [a GeoRSS feed of current major incidents](http://www.rfs.nsw.gov.au/feeds/majorIncidents.xml). As a developer, if you want bushfire data, you need to parse this feed to get the latest reports for current incidents. A [complication](https://github.com/andrewharvey/map.rfs/blob/master/rfs-major-incident-georss-to-geojson.pl#L3) can arise because the published GeoRSS is invalid - in GeoRSS items are only allowed to have 1 geometry, whereas sometimes the RFS GeoRSS feed includes several.

WIP