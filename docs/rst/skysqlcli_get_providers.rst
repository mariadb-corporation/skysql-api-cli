.. _skysqlcli_get_providers:

skysqlcli get providers
-----------------------

Retrieve list of cloud providers

Synopsis
~~~~~~~~


Queries a list of cloud providers supported by MariaDB SkySQL

::

  skysqlcli get providers [flags]

Options
~~~~~~~

::

  -h, --help   help for providers

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.skysql.net")
  -l, --limit int        Number of records to return (default 10)
      --mdbid string     URL for MariaDB ID (default "https://id.mariadb.com")

SEE ALSO
~~~~~~~~

* `skysqlcli get <skysqlcli_get.rst>`_ 	 - Query MariaDB SkySQL for resource information

