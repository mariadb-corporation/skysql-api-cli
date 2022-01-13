.. _skysqlcli_get_credentials:

skysqlcli get credentials
-------------------------

Retrieve default service credentials

Synopsis
~~~~~~~~


Queries for default credentials configured for a service. Specify a service using the service id (e.g. db00000000)

::

  skysqlcli get credentials [SERVICE] [flags]

Options
~~~~~~~

::

  -h, --help   help for credentials

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.skysql.net")
      --mdbid string     URL for MariaDB ID (default "https://id.mariadb.com")

SEE ALSO
~~~~~~~~

* `skysqlcli get <skysqlcli_get.rst>`_ 	 - Query MariaDB SkySQL for resource information

