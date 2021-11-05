.. _skysqlcli_get_tiers:

skysqlcli get tiers
-------------------

Retrieve list of MariaDB SkySQL account tiers

Synopsis
~~~~~~~~


Retrieves list of account tiers available for use with MariaDB SkySQL

::

  skysqlcli get tiers [flags]

Options
~~~~~~~

::

  -h, --help   help for tiers

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.dev.gcp.mariadb.net")
  -l, --limit int        Number of records to return (default 10)
      --mdbid string     URL for MariaDB ID (default "https://id-dev.mariadb.com")

SEE ALSO
~~~~~~~~

* `skysqlcli get <skysqlcli_get.rst>`_ 	 - Query MariaDB SkySQL for resource information

