## skysqlcli update service

Update an existing service

### Synopsis

Submits request to MariaDB SkySQL to update an existing service. Specify a service using the service id (e.g. db00000000)

```
skysqlcli update service [SERVICE] [flags]
```

### Options

```
  -h, --help          help for service
  -n, --name string   Name used to identify the service
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

