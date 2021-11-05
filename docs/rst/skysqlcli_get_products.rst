.. _skysqlcli_get_products:

skysqlcli get products
----------------------

Retrieve MariaDB SkySQL product information

Synopsis
~~~~~~~~


Queries information for product offerings from MariaDB SkySQL

::

  skysqlcli get products [flags]

Options
~~~~~~~

::

  -h, --help   help for products

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

