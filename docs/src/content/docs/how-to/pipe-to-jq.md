---
title: Multiple Locations with OpenAQ CLI & jq
description: A reference page in my new Starlight docs site.
---


#### Objective
This guide explains how to use the OpenAQ CLI in conjunction with `jq` to fetch air quality measurements from multiple locations over a specified date range and aggregate this data into a single JSON array.

#### Prerequisites
- `OpenAQ CLI` installed and configured
- `jq` installed on your system
- Basic understanding of shell scripting and JSON

#### Step-by-Step Guide

1. **Define Date Range for Measurements**
   - Set the start and end dates for the data collection
   - Example:
     ```bash
     from_date="2023-08-01"
     to_date="2023-08-02"
     ```

2. **Fetch List of Location IDs**
   - Use `openaq locations list` to get locations, and `jq` to extract their IDs. The `-r` flag in `jq` provides raw output, stripping quotes from JSON strings, making it suitable for shell scripts
   - Example for Great Britain (ISO code 'GB') and limiting results to 10 locations:
     ```bash
     location_ids=$(openaq locations list --iso GB --limit 10 --json | jq '.results[].id' -r)
     ```

3. **Initialize an Empty JSON Array**
   - Create an array to hold all combined measurements.
   - Example:
     ```bash
     all_measurements="[]"
     ```

4. **Loop Through Location IDs and Fetch Measurements**
   - For each ID, fetch measurements using `openaq measurements list`. `jq` filters and formats these measurements from the JSON output
   - Example for PM2.5 (parameter ID `2`):
     ```bash
     measurements=$(openaq measurements list $id --from ${from_date} --to ${to_date} --parameters 2 --json | jq '.results[]')
     ```

5. **Check and Format Measurements**
   - Convert the raw measurements to a properly formatted JSON array.
   - Example:
     ```bash
     if [[ -n $measurements ]]; then
         formatted_measurements=$(echo "$measurements" | jq -s '.')
     fi
     ```

6. **Aggregate Measurements**
   - Append each set of formatted measurements to the main array using `jq`. The `--argjson` flag helps in passing JSON arrays as arguments.
   - Example:
     ```bash
     all_measurements=$(jq -n --argjson current "$all_measurements" --argjson new "$formatted_measurements" '$current + $new')
     ```

7. **Output the Combined JSON Array**
   - Output the final aggregated data as a JSON array
   - Example:
     ```bash
     echo "$all_measurements" | jq '.'
     ```

8. **Run in a Shell Script**
  - Combine the commands from the previous steps into a shell script
  - Save the script as `measurements.sh` and make it executable with `chmod +x measurements.sh`
  - Run the script with `./measurements.sh`
  - Script:


```bash
#!/bin/bash

from_date="2023-08-01"
to_date="2023-08-02"

# Fetch list of location IDs
location_ids=$(openaq locations list --iso GB --limit 10 --json | jq '.results[].id' -r)

# Initialize an empty JSON array
all_measurements="[]"

# Loop through each location ID and fetch measurements
for id in $location_ids; do
    # Fetch measurements
    measurements=$(openaq measurements list $id --from ${from_date} --to ${to_date} --parameters 2 --json | jq '.results[]')

    # Check if measurements are not empty
    if [[ -n $measurements ]]; then
        # Convert the measurements to a proper JSON array
        formatted_measurements=$(echo "$measurements" | jq -s '.')

        # Append to all_measurements
        all_measurements=$(jq -n --argjson current "$all_measurements" --argjson new "$formatted_measurements" '$current + $new')
    fi
done

# Output final combined JSON array
echo "$all_measurements" | jq '.'
```

#### Output
This script outputs a JSON array of measurements from multiple locations over a specified date range. The output can be piped to other commands or saved to a file. It uses `jq` to filter and format the data, and the OpenAQ CLI to fetch the data.