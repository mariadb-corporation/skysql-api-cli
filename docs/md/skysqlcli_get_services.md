## skysqlcli get services

Retrieve service information

### Synopsis

Queries for information about deployed service resources in SkySQL. Specify a service using the service id (e.g. db00000000)

```
skysqlcli get services [SERVICE] [flags]
```

### Options

```
  -h, --help          help for services
  -n, --name string   Search string to match any services containing the name
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

