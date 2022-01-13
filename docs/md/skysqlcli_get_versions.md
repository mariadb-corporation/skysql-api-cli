## skysqlcli get versions

Retrieve list of MariaDB SkySQL service versions

### Synopsis

Retrieves list of service versions available for use with MariaDB SkySQL

```
skysqlcli get versions [flags]
```

### Options

```
  -h, --help         help for versions
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

