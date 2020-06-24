Listed all available configuration parameters here

Environment Variables
| CONFIG NAME | Description | Accepted Values |
| --------- | ----------- | --------------- |
| CONFIG_TYPE | localfile or remote |
| CONFIG_FORMAT | supported formats JSON, TOML, YAML, HCL, envfile and Java properties config |
| ENC_SECRET_KEY | Secret Key for Encryption. To Be Set as Environment Variable.Ignore Setting this if you dont want the security filters & utilities bundled in this library |
| ENC_TOKEN_EXPIRY | JWT Token Expiry Timeout. Value in Milliseconds. To Be Set as Environment Variable.Ignore Setting this if you want the default of 1 minute |
| ENV_CONFIG_PREFIX | Prefix for environment variables to make available in app |
| REMOTE_CONFIG_SECURE | Secure Remote config fetch enable/disable|Y/N| | Y/N |
| REMOTE_CONFIG_STORE | Type of Remote Config store | etcd/consul/firebase |
| REMOTE_CONFIG_HOST | Host address of remote config store |
| REMOTE_CONFIG_KEY | Remote config key |
| REMOTE_CONFIG_KEYRING | Remote config store security keyring |

Config Store - Reads Configuration Data from Configured Config Store, File or Remote. Define key in format compliant with format of your choice

- Database Properties

| CONFIG NAME    | Description                                                          | Accepted Values |
| -------------- | -------------------------------------------------------------------- | --------------- |
| db.enable      | Enable or Disable Bootstrap InBuilt Database Pool                    | Y/N             |
| db.driver      | SQL Database Driver.Currently Supporting pgx and go-sql-driver/mysql | pgx/mysql       |
| db.url         | SQL Database Server URL                                              |                 |
| db.max-open    | Max Open Database Connections. Default: 5                            |                 |
| db.max-idle    | Max Idle Database Connections. Default: 5                            |                 |
| db.max-timeout | Max Timeout Duration For Database Connections. Default: 2000ms       |                 |

- Logging Properties

  | CONFIG NAME               | Description                                                               | Accepted Values                            |
  | ------------------------- | ------------------------------------------------------------------------- | ------------------------------------------ |
  | logging.level             | Logging Level                                                             | info/error/debug/warn/fatal                |
  | logging.formatter.name    | LogData Formatting                                                        | json/text                                  |
  | logging.formatter.options | [Formatting Options](https://github.com/heralight/logrus_mate#formatters) |                                            |
  | logging.hooks             | [Hooks for Emitting Data](https://github.com/heralight/logrus_mate#hooks) | Currently Supported: graylog/logstash/file |
