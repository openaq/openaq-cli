---
title: Tutorial
description: A guide in my new Starlight docs site.
---

This tutorial provides a practical example of using OpenAQ CLI. We will identify locations within a specified bounding box for El Paso, TX. We can then choose a location and use its ID to retrieve measurements and export the data to a csv file.

### Find Locations using a Bounding Box
**Determine the Bounding Box for El Paso, TX:**
   - We will efine a bounding box that covers the El Paso area. The bounding box is defined in the format `minx,miny,maxx,maxy`. For example, `-106.6, 31.6, -106.2, 31.9`.

**Construct the Command:**
   - Use the `openaq locations list` command with `--bbox` for the bounding box:  
   - 
     ```
     openaq locations list --bbox -106.6,31.6,-106.2,31.9 
     ```

**Example response:**
   - Note down the `LOCATIONS_ID` column. Lets use ID `1089` to get measurements for the Ivanhoe C414 location.  


   ```bash
+--------------+-------------------------------------------------+--------------+-------------+--------------------------+-----------+-------------+
| LOCATIONS_ID | NAME                                            | COUNTRIES_ID | COUNTRY_ISO | COUNTRY_NAME             | LATITUDE  | LONGITUDE   |
+--------------+-------------------------------------------------+--------------+-------------+--------------------------+-----------+-------------+
| 1089         | Ivanhoe C414                                    | 13           | US          | United States of America | 31.786400 | -106.324200 |
| 1296         | Ascarate Park Southe                            | 13           | US          | United States of America | 31.746700 | -106.402802 |
| 1812         | Chamizal C41                                    | 13           | US          | United States of America | 31.765600 | -106.455000 |
   ```

### Get Measurements using a Location ID

**Construct the Command for Measurements:**
   - Use the `openaq measurements list` command with the chosen location ID. We can specify which pollutant we are interested in with the `parameters` flag. For instance, to get o3 measurements from August 1st, 2023, to August 2nd, 2023 and format as json:  
     ```
     openaq measurements list 1089 --from 2023-08-01 --to 2023-08-02 --parameters 10 --json
     ```

**Export Measurements as CSV:**
   - To format the data as a CSV, add the `--csv` flag and direct the output to a file:  
     ```
     openaq measurements list 1234 --from 2023-08-01 --to 2023-08-07 --parameters 10 --csv > ElPasoOzoneData.csv
     ```
   - This command saves the O3 measurements in a file named `ElPasoOzoneData.csv`


