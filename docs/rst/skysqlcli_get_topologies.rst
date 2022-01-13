.. _skysqlcli_get_topologies:

skysqlcli get topologies
------------------------

Retrieve list of MariaDB SkySQL service topologies

Synopsis
~~~~~~~~


Retrieves list of service topologies available for use with MariaDB SkySQL

::

  skysqlcli get topologies [flags]

Options
~~~~~~~

::

  -h, --help                  help for topologies
  -l, --limit int             Number of records to return. Can be used for paginating results in conjuntion with offset. (default 100)
  -o, --offset int            Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.
  -t, --service-type string   MariaDB SkySQL service-type used to filter list of topologies

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

