## skysqlcli get configurations

Retrieve stored service configurations

### Synopsis

Retrieves one or more custom service configurations owned by the user

```
skysqlcli get configurations [CONFIGURATION] [flags]
```

### Options

```
  -h, --help        help for configurations
  -l, --limit int   Number of records to return. Can be used for paginating results in conjuntion with offset. (default 100)
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

