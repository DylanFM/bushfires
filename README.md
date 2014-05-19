# Bushfires

_Part of a collection of tools used to provide an API to NSW bushfire data_: [Data collector](https://github.com/dylanfm/major-incidents-data), [Importer](https://github.com/DylanFM/incident-worker) and [GeoJSON API (this repo)](https://github.com/DylanFM/bushfires)

**NOTE:** This project is a work in progress. There is no public API offered yet.

The NSW Rural Fire Service publish [RSS feeds](http://www.rfs.nsw.gov.au/dsp_content.cfm?cat_id=1358) of incident data. In particular, they provide [a GeoRSS feed of current major incidents](http://www.rfs.nsw.gov.au/feeds/majorIncidents.xml). Developers seeking bushfire data must [parse](https://github.com/andrewharvey/map.rfs/blob/master/rfs-major-incident-georss-to-geojson.pl#L3) this feed to get the latest reports for current incidents.

The aim of this project is to provide more data than just the current incidents. It should provide access to historical data, all the reports for a given incident, as well as offering the data that the current RFS GeoRSS provides. On top of providing access to more data, this project aims to deliver it in a more developer-friendly format: [GeoJSON](http://geojson.org).

## Usage

This API is built using the [Go](http://golang.org) library [Tigertonic](http://github.com/rcrowley/go-tigertonic). It works with data stored in a Postres database. Refer to [the importer's README](https://github.com/DylanFM/incident-worker/blob/master/README.md) for information on setting up the database.

Once your database is set up, set your DATABASE_URL environment variable, e.g. `postgres://user:pass@localhost/database_name?sslmode=disable`.

To serve the API, run:

```
$ bushfires
```

There are additional config options available for this command, which can be viewed by running `bushfires --help`.

## API Documentation

Visit [the API docs](http://api.bushfir.es).

