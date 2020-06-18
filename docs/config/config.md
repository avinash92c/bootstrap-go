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
|CONFIG NAME|Type|Description|Accepted Values
|db.enable|Database Properties|Enable or Disable Bootstrap InBuilt Database Pool|Y/N
|
