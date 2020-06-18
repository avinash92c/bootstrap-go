Listed all available configuration parameters here

Environment Variables
| CONFIG NAME | Description | Accepted Values |
| --------- | ----------- | --------------- |
| CONFIG_TYPE | localfile or remote | 
| CONFIG_FORMAT | supported formats JSON, TOML, YAML, HCL, envfile and Java properties config |
| ENV_CONFIG_PREFIX | Prefix for environment variables to make available in app |
| REMOTE_CONFIG_SECURE | Secure Remote config fetch enable/disable|Y/N| | Y/N |
| REMOTE_CONFIG_STORE | Type of Remote Config store | etcd/consul/firebase |
| REMOTE_CONFIG_HOST | Host address of remote config store |
| REMOTE_CONFIG_KEY | Remote config key |
| REMOTE_CONFIG_KEYRING | Remote config store security keyring |

Config Store - Reads Configuration Data from Configured Config Store, File or Remote. Define key in format compliant with format of your choice
| Type | CONFIG NAME | Description | Accepted Values |
| --------- | ----------- | --------------- |
| Database Properties | db.enable | Enable or Disable Bootstrap InBuilt Database Pool | Y/N |
| Database Properties | db.driver | SQL Database Driver.Currently Supporting pgx and go-sql-driver/mysql | pgx/mysql |
| Database Properties | db.url | SQL Database Server URL |  |
| Database Properties | db.max-open | Max Open Database Connections. Default: 5 |  |
| Database Properties | db.max-idle | Max Idle Database Connections. Default: 5 |  |
| Database Properties | db.max-timeout | Max Timeout Duration For Database Connections. Default: 2000ms |  |