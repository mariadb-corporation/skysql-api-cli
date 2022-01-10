## skysqlcli delete allowed-address

Delete an allowed address from a service

### Synopsis

Deletes an allowed address for user on service in MariaDB SkySQL.

```
skysqlcli delete allowed-address [SERVICE] [ALLOWED-ADDRESS] [flags]
```

### Options

```
  -h, --help   help for allowed-address
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

