---
title: flags
description: A reference page in my new Starlight docs site.
prev: false

sidebar:
  order: 1
---


## Global flags

`--config` `-c` — Manually set the path to a configuration TOML file, default is $HOME/.openaq.toml

`--help` `-h` — returns helps for any command

## Resource flags 

`--json` - Returns data as JSON instead of the default table view.

`--pretty` - Only used in combination with `--json`. Provides a "pretty" indented and syntax highlighted view of the JSON output

`--csv` - Returns the result as a csv (comma separated values).


_the following only work on resource `list` calls_

`--limit` - takes an integer to limit the number of results to return.

`--page` - takes an integer to set the page number for API pagination