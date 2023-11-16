---
title: Configuration file
description: A reference page in my new Starlight docs site.

prev: false
---
The OpenAQ CLI requires a configuration file to store and access certain variables on the local file system. The first time `openaq` is run a file named `.openaq.toml` will be created in the `$HOME` directory of your system.


> The configuration file specification is not yet feature complete, changes to the specification will likely occur in future releases.

`.openaq.toml` format:

```toml
api-key = ''

[defaults]
format = 'json'
pretty = false
```