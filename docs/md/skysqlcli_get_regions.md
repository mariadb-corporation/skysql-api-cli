## skysqlcli get regions

Retrieve list of provider regions

### Synopsis

Retrieves list of provider regions available for use with MariaDB SkySQL

```
skysqlcli get regions [flags]
```

### Options

```
  -h, --help              help for regions
      --provider string   MariaDB SkySQL provider to query for regions
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

