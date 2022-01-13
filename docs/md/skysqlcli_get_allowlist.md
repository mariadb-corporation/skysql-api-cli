## skysqlcli get allowlist

Retrieve list of allowed IPs for a service

### Synopsis

Queries for list of allowed IP addresses for a specific service

```
skysqlcli get allowlist [SERVICE] [flags]
```

### Options

```
  -h, --help         help for allowlist
  -l, --limit int    Number of records to return. Can be used for paginating results in conjuntion with offset. (default 100)
  -o, --offset int   Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.
```

### Options inherited from parent commands

```
      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.skysql.net")
      --mdbid string     URL for MariaDB ID (default "https://id.mariadb.com")
```

### SEE ALSO

* [skysqlcli get](skysqlcli_get.md)	 - Query MariaDB SkySQL for resource information

