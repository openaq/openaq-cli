---
title: locations
description: A reference page in my new Starlight docs site.
---


```sh
openaq locations list 
```

Provides a list of locations.

__flags__

`--countries` - filter locations by one or more comma-delimited `countries ID` e.g. 1,2,3

`--providers` - filter locations by one or more comma-delimited `providers ID` e.g. 1,2,3

`--iso` - Filter the results to a specific country using ISO 3166-1 alpha-2 code

`--radius` - The radius in meters to search around the `--coordinates` center point. Must be used with `--coordinates`

`--coordinates` - The center point coordinates for the radius search in form latitude,longitude i.e. y,x. Must be used with `--radius`

`--bbox` - Filter results to those contained within a spatial bounding box, in form minx,miny,maxx,maxy. Cannot be used with `--radius`/`--coordinates`

---

```sh
openaq locations get [locationsID]
```

Provides a single location given a locationsID

