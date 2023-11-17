---
title: Examples
description: A guide in my new Starlight docs site.
---

Get locations withing the radius of 5000 meters from the center point of 40.7128,-74.0060.
```sh
openaq locations list --coordinates 40.7128,-74.0060 --radius 5000
```

-------
Get data for location 2178.


```sh   
openaq locations get 2178
```

------
Get measurements for location 2178 between 2023-08-01 and 2023-08-07 and return the results in csv format.

```sh
openaq measurements list 2178 --from 2023-08-01 --to 2023-08-07 --csv
```

-------
Get measurements for location 2178 between 2023-08-01 and 2023-08-07 and limit to only PM2.5 return the results in json format.

```sh
openaq measurements list 2178 --from 2023-08-01 --to 2023-08-07 --parameters 2 --json
```

-------
Get locations only from provider_id 166 (Clarity) 

```sh
openaq locations list --provider 166
```

-----