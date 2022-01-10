.. _skysqlcli_get_versions:

skysqlcli get versions
----------------------

Retrieve list of MariaDB SkySQL service versions

Synopsis
~~~~~~~~


Retrieves list of service versions available for use with MariaDB SkySQL

::

  skysqlcli get versions [flags]

Options
~~~~~~~

::

  -h, --help   help for versions

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

