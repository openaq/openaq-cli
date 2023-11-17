---
title: How to find locations
description: A reference page in my new Starlight docs site.
---

### Find Locations by ISO Code
**Determine the ISO Code:**
   - Identify the ISO 3166-1 alpha-2 code of the country for which you want to find locations. For instance, `US` for the United States.

**Construct the Command:**
   - Use the `openaq locations list` command with the `--iso` flag followed by the ISO code:
     ```
     openaq locations list --iso US
     ```
   - This command will list all locations in the United States.

### Find Locations by Coordinates and Radius
**Determine the Coordinates and Radius:**
   - Specify the center point coordinates in the form `latitude,longitude` and the radius in meters for the search. For example, coordinates `40.7128,-74.0060` with a 5000-meter radius.

**Construct the Command:**
   - Use the `openaq locations list` command with the `--coordinates` and `--radius` flags:
     ```
     openaq locations list --coordinates 40.7128,-74.0060 --radius 5000
     ```
   - This command will list all locations within a 5000-meter radius of the given coordinates.

### Find Locations by Bounding Box
**Determine the Bounding Box:**
   - Define the bounding box in the format `minx,miny,maxx,maxy`. For example, a box around New York City might be `-74.2599,40.4774,-73.7004,40.9176`.

**Construct the Command:**
   - Use the `openaq locations list` command with the `--bbox` flag followed by the bounding box coordinates:
     ```
     openaq locations list --bbox -74.2599,40.4774,-73.7004,40.9176
     ```
   - This command will list all locations within the specified bounding box.

