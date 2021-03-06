.. _skysqlcli_get_sizes:

skysqlcli get sizes
-------------------

Retrieve list of machine sizes

Synopsis
~~~~~~~~


Retrieves list of machine sizes available for use with MariaDB SkySQL

::

  skysqlcli get sizes [flags]

Options
~~~~~~~

::

  -h, --help                  help for sizes
  -l, --limit int             Number of records to return. Can be used for paginating results in conjuntion with offset. (default 100)
  -o, --offset int            Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.
      --provider string       MariaDB SkySQL provider to query for machine sizes
      --service-type string   MariaDB SkySQL service-type to query for machine sizes (default "t")
      --tier string           MariaDB SkySQL tier to query for machine sizes

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

