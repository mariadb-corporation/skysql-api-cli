.. _skysqlcli_get_configurations:

skysqlcli get configurations
----------------------------

Retrieve stored service configurations

Synopsis
~~~~~~~~


Retrieves one or more custom service configurations owned by the user

::

  skysqlcli get configurations [CONFIGURATION] [flags]

Options
~~~~~~~

::

  -h, --help   help for configurations

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

