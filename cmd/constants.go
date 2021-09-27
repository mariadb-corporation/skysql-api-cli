package cmd

// Identifiers
const (
	CLI_NAME     = "skysqlcli"
	PROJECT_NAME = "skysql"
)

// URLs
const (
	SKYSQL_API = "https://api.dev.gcp.mariadb.net"
	MARIADB_ID = "https://id-dev.mariadb.com"
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
	ACCOUNT         = "account"
	ACCOUNTS        = "accounts"
	ALLOWLIST       = "allowlist"
	ALLOWED_ADDRESS = "allowed-address"
	COMMENT         = "comment"
	CONFIGURATION   = "configuration"
	CONFIGURATIONS  = "configurations"
	DATABASE        = "database"
	DATABASES       = "databases"
	IP_ADDRESS      = "ip-address"
	PRODUCT         = "product"
	PRODUCTS        = "products"
	PROVIDER        = "provider"
	PROVIDERS       = "providers"
	QUOTA           = "quota"
	QUOTAS          = "quotas"
	REGION          = "region"
	REGIONS         = "regions"
	SIZE            = "size"
	SIZES           = "sizes"
	STORAGE         = "storage"
	TIER            = "tier"
	TIERS           = "tiers"
	TOPOLOGIES      = "topologies"
	TOPOLOGY        = "topology"
	VERSION         = "version"
	VERSIONS        = "versions"
)

// hints
const (
	HINT_DB_ID = "Specify a database using the database id (e.g. db00000000)"
	HINT_LIMIT = "Maximum number of results to retrieve"
)

// defaults
const (
	DEFAULT_GET_LIMIT = 10

	DEFAULT_CREATE_DATABASE_RELEASE_VERSION = "MariaDB Enterprise Server 10.5.9-6"
	DEFAULT_CREATE_DATABASE_TOPOLOGY        = "Standalone"
	DEFAULT_CREATE_DATABASE_SIZE            = "Sky-2x4"
	DEFAULT_CREATE_DATABASE_TXSTORAGE       = "100"
	DEFAULT_CREATE_DATABASE_MAXSCALE_CONFIG = ""
	DEFAULT_CREATE_DATABASE_NAME            = ""
	DEFAULT_CREATE_DATABASE_REGION          = ""
	DEFAULT_CREATE_DATABASE_REPL_REGION     = ""
	DEFAULT_CREATE_DATABASE_PROVIDER        = ""
	DEFAULT_CREATE_DATABASE_REPLICAS        = "0"
	DEFAULT_CREATE_DATABASE_MONITOR         = "false"
	DEFAULT_CREATE_DATABASE_MAXSCALE_PROXY  = "false"
	DEFAULT_CREATE_DATABASE_VOLUME_IOPS     = ""
	DEFAULT_CREATE_DATABASE_VOLUME_TYPE     = ""

	DEFAULT_UPDATE_DATABASE_NAME = ""
)

// Flags
const (
	CONFIG_JSON     = "config-json"
	LIMIT           = "limit"
	LIMIT_SHORTHAND = "l"
	MAXSCALE_CONFIG = "maxscale-config"
	MAXSCALE_PROXY  = "maxscale-proxy"
	MONITOR         = "monitor"
	NAME            = "name"
	RELEASE_VERSION = "release-version"
	REPLICAS        = "replicas"
	REPL_REGION     = "repl-region"
	VOLUME_IOPS     = "volume-iops"
	VOLUME_TYPE     = "volume-type"
)
