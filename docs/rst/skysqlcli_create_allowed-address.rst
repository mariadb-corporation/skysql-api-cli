.. _skysqlcli_create_allowed-address:

skysqlcli create allowed-address
--------------------------------

Add a new allowed address to a service

Synopsis
~~~~~~~~


Adds a new allowed address for service in MariaDB SkySQL.

::

  skysqlcli create allowed-address [SERVICE] [IP-ADDRESS] [flags]

Options
~~~~~~~

::

      --comment string   Additional comment to help identify address
  -h, --help             help for allowed-address

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.skysql.net")
      --mdbid string     URL for MariaDB ID (default "https://id.mariadb.com")

SEE ALSO
~~~~~~~~

* `skysqlcli create <skysqlcli_create.rst>`_ 	 - Create a resource in MariaDB SkySQL

