.. _skysqlcli_update_status:

skysqlcli update status
-----------------------

Update status for service

Synopsis
~~~~~~~~


Updates status for service belonging user in MariaDB SkySQL.

::

  skysqlcli update status [SERVICE] [Start|Stop] [flags]

Options
~~~~~~~

::

  -h, --help   help for status

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.skysql.net")
      --mdbid string     URL for MariaDB ID (default "https://id.mariadb.com")

SEE ALSO
~~~~~~~~

* `skysqlcli update <skysqlcli_update.rst>`_ 	 - Update a resource in MariaDB SkySQL

