---
title: measurements
description: A reference page in my new Starlight docs site.
---

```sh
openaq measurements list [locationsID]
```

Provides a list of measurements for a given locationsID


__flags__

`--to` - a date in form YYYY-MM-DD to filter measurements as the end date of the period e.g. 2023-08-23

`--from` -  a date in form YYYY-MM-DD to filter measurements as the start date of the period e.g. 2023-08-23

`--parameters` - filter measurements by one or more comma-delimited `parameters ID` e.g. 1,2,3

`--mini` - provides a miniature/simplified version of the measurements resource including only the parameter name, datetime in UTC, datetime in local, measurement period, and value 
