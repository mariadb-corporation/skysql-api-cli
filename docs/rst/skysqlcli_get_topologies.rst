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
  -t, --service-type string   MariaDB SkySQL service-type used to filter list of topologies

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

