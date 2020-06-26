BootStrap-GO Usage Guide

- Configuration
  Currently two modes of Configuration Sources are Accepted.

  - File
  - Remote

    Supported Configuration Formats are

    - JSON, TOML, YAML, HCL, envfile and Java properties config files

- File Config
  The default config path for reading config files is "./config". A folder named config in your application root dir.
  Alternatively if you wish to specify a custom location to read configurations,
  you could do so with the command line flag "configpath"
  e.g., app.exe -configpath="/path/to/config"
