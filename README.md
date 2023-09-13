
# OpenAQ API Command Line Client

A command line utility to interact with the OpenAQ API.

![parameters](https://github.com/openaq/openaq-cli/assets/8487728/07a860ee-31cf-4ee3-b52c-6146c9617536)

## Installation

> The OpenAQ CLI tool is still a *work in progress* and may be unstable until a 1.0 release. 

Compiled executables for Windows, Mac, Linux are available for download in the releases page: 
https://github.com/openaq/openaq-cli/releases/

After downloading place the exectuable where you keep exectuables and update you system `$PATH` variable to make the executable discoverable by your shell.

---

Alternatively you can install with Golang > 1.18 with:

```
go install github.com/openaq/openaq-cli
```



## Usage

> Note: An OpenAQ API Key is required to use the OpenAQ CLI. Registering for an API Key is free and only requires a valid email address. Register at https://api.openaq.org/register


To set your API Key run: 

```sh
openaq configure api-key set my-super-secret-openaq-api-key-1234-5678
```

replacing `my-super-secret-openaq-api-key-1234-5678` with your API Key


### Global flags

`--config` `-c` — Manually set a configuration TOML file, default is $HOME/.openaq.toml

`--help` `-h` — returns helps for any command

### Resource flags 

`--json` - Returns data as JSON instead of the default table view.

`--pretty` - Only used in combination with `--json`. Provides a "pretty" indented and syntax highlighted view of the JSON output

`--csv` - Returns the result as a csv (comma separated values).


_the following only work on resource `list` calls_

`--limit` - takes an integer to limit the number of results to return.

`--page` - takes an integer to set the page number for API pagination

### Commands

about 
---

```sh
openaq about
```

Provides a description of the OpenAQ CLI.


configure 
---
__api-key__

```sh
openaq configure api-key set [api-key]
```

countries
---
```sh
openaq countries list 
```

Provides a list of countries.


__flags__

`--mini` -  provides a miniature/simplified version of the countries resource including only the countries ID, ISO code, name, and parameters

---


```sh
openaq countries get [countriesID]
```

Provides a single country given a countriesID


help
---
```
openaq help
```

Provides a help guide with commands


locations
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


manufacturers
---

Coming soon

measurements
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


owners
---
Coming soon

parameters
---

```sh
openaq parameters list 
```
Provides a list of parameters.


__flags__

`--type` - filter parameters by `parameterType` either `pollutant` or `meteorological`


---


```sh
openaq parameters get [parametersID]
```
Provides a single parameter for a given parametersID


version 
---
```
openaq version
```

prints the version of the OpenAQ CLI


## Examples

Get measurements for location 2178 between 2023-08-01 and 2023-08-07 and return the results as a csv.

```sh
openaq measurements list 2178 --from 2023-08-01 --to 2023-08-07 --csv
```

Get measurements for location 2178 between 2023-08-01 and 2023-08-07 and limit to only PM2.5 return the results as a csv.

```sh
openaq measurements list 2178 --from 2023-08-01 --to 2023-08-07 --parameters 2 --csv
```

Get locations only from provider_id 166 (Clarity) 

```sh
openaq locations list --provider 166
```

### Configuration file

The OpenAQ CLI requires a configuration file to store and access certain variables on the local file system. The first time `openaq` is run a file named `.openaq.toml` will be created in the `$HOME` directory of your system.


> The configuration file specification is not yet feature complete, changes to the specification will likely occur in future releases.

`.openaq.toml` format:

```toml
api-key = ''
```


