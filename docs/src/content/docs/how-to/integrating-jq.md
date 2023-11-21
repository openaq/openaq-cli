---
title: Processing data with jq
description: A reference page in my new Starlight docs site.
---
### Overview

This guide demonstrates how to use jq to manipulate and analyze data from JSON output obtained via the OpenAQ CLI tool. jq is a command-line JSON processor that allows you to extract and manipulate data from JSON files. It is available for Linux, macOS, and Windows.

### Prerequisites
- `OpenAQ CLI` installed and configured.
- `jq` already installed on your system.

### Step 1: Fetching Data with OpenAQ CLI

Fetch air quality measurements, for example:

```bash
openaq measurements list 2178 --format=json > measurements.json
```
This example command saves air quality measurements from Albuquerque, NM to measurements.json.

### Step 2: Basic Usage of jq
To view the JSON data in a readable format:

```bash
jq '.' measurements.json
```
### Step 3: Extracting Specific Data with jq
To extract specific information like parameter, units, and measurement value:

```bash
jq '.results[] | {parameter: .parameter.name, value: .value, unit: .parameter.units}' measurements.json
```
### Step 4: Counting Objects in the Data
To count the number of measurement objects in the JSON file:

```bash
jq '.results | length' measurements.json
```
This jq command calculates the number of elements in the results array, effectively counting the measurement entries.

### Step 5: Advanced Data Processing
For more complex tasks, such as filtering measurements with values greater than a certain threshold:

```bash
jq '.results[] | select(.value > 10)' measurements.json

```
jq is a versatile tool that enhances your ability to process and analyze JSON data from the OpenAQ platform. This guide covers fundamental operations, from viewing data to counting objects and applying filters.

For additional information on jq's capabilities, refer to the [jq manual](https://stedolan.github.io/jq/manual/).

