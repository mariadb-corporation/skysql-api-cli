.. _skysqlcli_get_services:

skysqlcli get services
----------------------

Retrieve service information

Synopsis
~~~~~~~~


Queries for information about deployed service resources in SkySQL. Specify a service using the service id (e.g. db00000000)

::

  skysqlcli get services [SERVICE] [flags]

Options
~~~~~~~

::

  -h, --help          help for services
  -l, --limit int     Number of records to return. Can be used for paginating results in conjuntion with offset. (default 100)
  -n, --name string   Search string to match any services containing the name
  -o, --offset int    Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.

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

