package cmd

// Identifiers
const (
	CLI_NAME     = "skysqlcli"
	PROJECT_NAME = "skysql"
)

// URLs
const (
	SKYSQL_API = "https://api.skysql.net"
	MARIADB_ID = "https://id.mariadb.com"
)

// Actions
const (
	GET    = "get"
	CREATE = "create"
	DELETE = "delete"
	UPDATE = "update"
)

// Entities
const (
	ACCOUNT          = "account"
	ACCOUNTS         = "accounts"
	ALLOWLIST        = "allowlist"
	ALLOWLIST_STATUS = "allowlist-status"
	ALLOWED_ADDRESS  = "allowed-address"
	COMMENT          = "comment"
	CONFIGURATION    = "configuration"
	CONFIGURATIONS   = "configurations"
	CREDENTIALS      = "credentials"
	SERVICE          = "service"
	SERVICES         = "services"
	IP_ADDRESS       = "ip-address"
	MACHINE          = "machine"
	SERVICE_TYPE     = "service-type"
	SERVICE_TYPES    = "service-types"
	PROVIDER         = "provider"
	PROVIDERS        = "providers"
	QUOTA            = "quota"
	QUOTAS           = "quotas"
	REGION           = "region"
	REGIONS          = "regions"
	SIZE             = "size"
	SIZES            = "sizes"
	STATUS           = "status"
	SSL_TLS          = "ssl-tls"
	STORAGE          = "storage"
	TIER             = "tier"
	TIERS            = "tiers"
	TOPOLOGIES       = "topologies"
	TOPOLOGY         = "topology"
	VERSION          = "version"
	VERSIONS         = "versions"
)

// hints
const (
	HINT_SVC_ID = "Specify a service using the service id (e.g. db00000000)"
	HINT_LIMIT  = "Maximum number of results to retrieve"
)

// defaults
const (
	DEFAULT_GET_LIMIT  = 100
	DEFAULT_GET_OFFSET = 0

	DEFAULT_CREATE_SERVICE_RELEASE_VERSION = ""
	DEFAULT_CREATE_SERVICE_TOPOLOGY        = "Single Node Transactions"
	DEFAULT_CREATE_SERVICE_SIZE            = "Sky-2x4"
	DEFAULT_CREATE_SERVICE_TXSTORAGE       = "100"
	DEFAULT_CREATE_SERVICE_MAXSCALE_CONFIG = ""
	DEFAULT_CREATE_SERVICE_NAME            = ""
	DEFAULT_CREATE_SERVICE_REGION          = ""
	DEFAULT_CREATE_SERVICE_REPL_REGION     = ""
	DEFAULT_CREATE_SERVICE_PROVIDER        = ""
	DEFAULT_CREATE_SERVICE_REPLICAS        = "0"
	DEFAULT_CREATE_SERVICE_MONITOR         = "false"
	DEFAULT_CREATE_SERVICE_MAXSCALE_PROXY  = "false"
	DEFAULT_CREATE_SERVICE_VOLUME_IOPS     = ""
	DEFAULT_CREATE_SERVICE_VOLUME_TYPE     = ""
	DEFAULT_CREATE_SERVICE_TIER            = ""
	DEFAULT_CREATE_SERVICE_SSL_TLS         = "Enabled"

	DEFAULT_UPDATE_SERVICE_NAME = ""

	DEFAULT_CREATE_CONFIGURATION_TOPOLOGY    = "Read Scalable Transactions"
	DEFAULT_CREATE_CONFIGURATION_NAME        = "HA"
	DEFAULT_CREATE_CONFIGURATION_CONFIG_JSON = `
		{
			"variables": {
				"interactive_timeout": {
					"type": "number",
					"value": "300",
					"requires_restart": true,
					"regex": ""
				}
			}
		}
	`
)

// Flags
const (
	CONFIG_JSON      = "config-json"
	LIMIT            = "limit"
	LIMIT_SHORTHAND  = "l"
	MAXSCALE_CONFIG  = "maxscale-config"
	MAXSCALE_PROXY   = "maxscale-proxy"
	MONITOR          = "monitor"
	NAME             = "name"
	OFFSET           = "offset"
	OFFSET_SHORTHAND = "o"
	RELEASE_VERSION  = "release-version"
	REPLICAS         = "replicas"
	REPL_REGION      = "repl-region"
	VOLUME_IOPS      = "volume-iops"
	VOLUME_TYPE      = "volume-type"
)
