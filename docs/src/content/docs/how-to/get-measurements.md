---
title: How to download measurements
description: A reference page in my new Starlight docs site.
---

### Step 1: Fetch Measurements
**Determine Your Date Range:**
   - Decide on the start and end dates for which you want to fetch measurements. The date format should be in the format `YYYY-MM-DD` (Year-Month-Day). For example, `2023-08-01` represents the 1st of August, 2023.

**Construct the Command:**
   - Use the `openaq measurements list` command with the location ID (for example, 2178), and specify your date range and parameter IDs. For example, if you want measurements from August 1st, 2023, to August 7th, 2023, the command would look like this:
     ```
     openaq measurements list 2178 --from 2023-08-01 --to 2023-08-07 --parameters 2
     ```
   - Use the --parameters flag with the parameters ID you are interested in (e.g., 2 for PM2.5).

### Step 2: Export as CSV
**Add CSV Flag:**
   - To format the retrieved data as a CSV, append `--csv` to the end of your command:
     ```
     openaq measurements list 2178 --from 2023-08-01 --to 2023-08-07 --parameters 2 --csv
     ```

**Save Output to File:**
   - Direct the output to a CSV file by adding `> YourFileName.csv` at the end of the command:
     ```
     openaq measurements list 2178 --from 2023-08-01 --to 2023-08-07 --parameters 2 --csv > YourFileName.csv
     ```
   - Replace `YourFileName.csv` with your preferred file name for the CSV.

