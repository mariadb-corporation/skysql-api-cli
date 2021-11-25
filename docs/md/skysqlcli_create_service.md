## skysqlcli create service

Create a new service

### Synopsis

Submits request to MariaDB SkySQL to deploy a new service

```
skysqlcli create service [flags]
```

### Options

```
  -h, --help                     help for service
  -m, --maxscale-config string   Configurations for maxscale
      --maxscale-proxy string    Whether to set up a proxy for maxscale (default "false")
      --monitor string           Whether to deploy a monitoring cluster alongside the service (default "false")
  -n, --name string              Name used to identify the service
  -p, --provider string          Cloud provider to host the service
  -r, --region string            Geographic region to deploy the service
  -v, --release-version string   Release version to deploy
      --repl-region string       Replica region for the service
      --replicas string          Number of replicas to deploy (default "0")
  -s, --size string              Size of the nodes running the service (default "Sky-2x4")
      --ssl-tls string           Specify whether to use SSL/TLS encryption (default "Enabled")
  -g, --storage string           Size of the persistent storage disk (default "100")
      --tier string              Tier in which to provision service
  -t, --topology string          Service topology to select (default "Single Node Transactions")
      --volume-iops string       Amount of IOPS for the volume (e.g. 100). Required for Amazon AWS
      --volume-type string       Type of volume to use (e.g. io1, gp3). Required for Amazon AWS
```

### Options inherited from parent commands

```
      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.dev.gcp.mariadb.net")
      --mdbid string     URL for MariaDB ID (default "https://id-dev.mariadb.com")
```

### SEE ALSO

* [skysqlcli create](skysqlcli_create.md)	 - Create a resource in MariaDB SkySQL

