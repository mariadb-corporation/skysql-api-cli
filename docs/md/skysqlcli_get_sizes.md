## skysqlcli get sizes

Retrieve list of machine sizes

### Synopsis

Retrieves list of machine sizes available for use with MariaDB SkySQL

```
skysqlcli get sizes [flags]
```

### Options

```
  -h, --help              help for sizes
      --product string    MariaDB SkySQL product to query for machine sizes
      --provider string   MariaDB SkySQL provider to query for machine sizes
      --tier string       MariaDB SkySQL tier to query for machine sizes
```

### Options inherited from parent commands

```
      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.dev.gcp.mariadb.net")
  -l, --limit int        Number of records to return (default 10)
      --mdbid string     URL for MariaDB ID (default "https://id-dev.mariadb.com")
```

### SEE ALSO

* [skysqlcli get](skysqlcli_get.md)	 - Query MariaDB SkySQL for resource information

