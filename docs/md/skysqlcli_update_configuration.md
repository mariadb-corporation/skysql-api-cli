## skysqlcli update configuration

Update a configuration

### Synopsis

Updates a configuration for user in MariaDB SkySQL.

```
skysqlcli update configuration [CONFIGURATION] [flags]
```

### Options

```
  -j, --config-json string   JSON object containing configuration
  -h, --help                 help for configuration
  -n, --name string          Name used to identify the configuration
```

### Options inherited from parent commands

```
      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.dev.gcp.mariadb.net")
      --mdbid string     URL for MariaDB ID (default "https://id-dev.mariadb.com")
```

### SEE ALSO

* [skysqlcli update](skysqlcli_update.md)	 - Update a resource in MariaDB SkySQL
