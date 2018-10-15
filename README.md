This is a Go package which maps a (lat, long) to a timezone which
forked from [bradfitz/latlong](https://github.com/bradfitz/latlong).

Since the data source
[tz_world.zip](http://efele.net/maps/tz/world/tz_world.zip)
is no longer maintained, we use another OSM-based map
[timezone-boundary-builder](https://github.com/evansiroky/timezone-boundary-builder)
instead.

The z_gen_tables.go was generated by `2018d` [release](https://github.com/evansiroky/timezone-boundary-builder/releases)
[`timezones.shapefile.zip`](https://github.com/evansiroky/timezone-boundary-builder/releases/download/2018d/timezones.shapefile.zip).