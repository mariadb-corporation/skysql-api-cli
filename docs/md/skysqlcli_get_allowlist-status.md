## skysqlcli get allowlist-status

Retrieve status of allowlist for a service

### Synopsis

Queries for the provisioning status of an allowlist for a specific service

```
skysqlcli get allowlist-status [SERVICE] [flags]
```

### Options

```
  -h, --help   help for allowlist-status
```

### Options inherited from parent commands

```
      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.skysql.net")
  -l, --limit int        Number of records to return (default 10)
      --mdbid string     URL for MariaDB ID (default "https://id.mariadb.com")
```

### SEE ALSO

* [skysqlcli get](skysqlcli_get.md)	 - Query MariaDB SkySQL for resource information

