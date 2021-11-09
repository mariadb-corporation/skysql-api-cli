## skysqlcli get

Query MariaDB SkySQL for resource information

### Synopsis

Commands which retrieve data about resources deployed into MariaDB SkySQL

### Options

```
  -h, --help        help for get
  -l, --limit int   Number of records to return (default 10)
```

### Options inherited from parent commands

```
      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.dev.gcp.mariadb.net")
      --mdbid string     URL for MariaDB ID (default "https://id-dev.mariadb.com")
```

### SEE ALSO

* [skysqlcli](skysqlcli.md)	 - CLI client to interact with the SkySQL API
* [skysqlcli get allowlist](skysqlcli_get_allowlist.md)	 - Retrieve list of allowed IPs for a service
* [skysqlcli get configurations](skysqlcli_get_configurations.md)	 - Retrieve stored service configurations
* [skysqlcli get products](skysqlcli_get_products.md)	 - Retrieve MariaDB SkySQL product information
* [skysqlcli get providers](skysqlcli_get_providers.md)	 - Retrieve list of cloud providers
* [skysqlcli get quotas](skysqlcli_get_quotas.md)	 - Retrieve quota information
* [skysqlcli get regions](skysqlcli_get_regions.md)	 - Retrieve list of provider regions
* [skysqlcli get services](skysqlcli_get_services.md)	 - Retrieve service information
* [skysqlcli get sizes](skysqlcli_get_sizes.md)	 - Retrieve list of machine sizes
* [skysqlcli get status](skysqlcli_get_status.md)	 - Get current status for a service
* [skysqlcli get tiers](skysqlcli_get_tiers.md)	 - Retrieve list of MariaDB SkySQL account tiers
* [skysqlcli get topologies](skysqlcli_get_topologies.md)	 - Retrieve list of MariaDB SkySQL service topologies
* [skysqlcli get versions](skysqlcli_get_versions.md)	 - Retrieve list of MariaDB SkySQL service versions
