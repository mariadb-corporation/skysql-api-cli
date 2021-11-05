.. _skysqlcli_create_configuration:

skysqlcli create configuration
------------------------------

Create a new configuration

Synopsis
~~~~~~~~


Creates a new configuration for user in MariaDB SkySQL.

::

  skysqlcli create configuration [flags]

Options
~~~~~~~

::

  -j, --config-json string   JSON object containing configuration (default "\n\t\t{\n\t\t\t\"variables\": {\n\t\t\t\t\"interactive_timeout\": {\n\t\t\t\t\t\"type\": \"number\",\n\t\t\t\t\t\"value\": \"300\",\n\t\t\t\t\t\"requires_restart\": true,\n\t\t\t\t\t\"regex\": \"\"\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t")
  -h, --help                 help for configuration
  -n, --name string          Name used to identify the configuration (default "HA")
  -t, --topology string      Configuration topology to select (default "Primary/Replica")

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.dev.gcp.mariadb.net")
      --mdbid string     URL for MariaDB ID (default "https://id-dev.mariadb.com")

SEE ALSO
~~~~~~~~

* `skysqlcli create <skysqlcli_create.rst>`_ 	 - Create a resource in MariaDB SkySQL

