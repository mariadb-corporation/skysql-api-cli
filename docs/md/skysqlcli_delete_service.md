## skysqlcli delete service

Delete an existing service

### Synopsis

Submits request to MariaDB SkySQL to delete an existing service. Specify a service using the service id (e.g. db00000000)

```
skysqlcli delete service [SERVICE] [flags]
```

### Options

```
  -h, --help   help for service
```

### Options inherited from parent commands

```
      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.skysql.net")
      --mdbid string     URL for MariaDB ID (default "https://id.mariadb.com")
```

### SEE ALSO

* [skysqlcli delete](skysqlcli_delete.md)	 - Delete a resource in MariaDB SkySQL

