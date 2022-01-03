.. _skysqlcli_get:

skysqlcli get
-------------

Query MariaDB SkySQL for resource information

Synopsis
~~~~~~~~


Commands which retrieve data about resources deployed into MariaDB SkySQL

Options
~~~~~~~

::

  -h, --help        help for get
  -l, --limit int   Number of records to return (default 10)

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --api-key string   Long-lived JWT issued from MariaDB ID
  -c, --config string    config file (default $HOME/.skysqlcli.yaml)
      --host string      URL for the SkySQL API (default "https://api.dev.gcp.mariadb.net")
      --mdbid string     URL for MariaDB ID (default "https://id-dev.mariadb.com")

SEE ALSO
~~~~~~~~

* `skysqlcli <skysqlcli.rst>`_ 	 - CLI client to interact with the SkySQL API
* `skysqlcli get allowlist <skysqlcli_get_allowlist.rst>`_ 	 - Retrieve list of allowed IPs for a service
* `skysqlcli get allowlist-status <skysqlcli_get_allowlist-status.rst>`_ 	 - Retrieve status of allowlist for a service
* `skysqlcli get configurations <skysqlcli_get_configurations.rst>`_ 	 - Retrieve stored service configurations
* `skysqlcli get credentials <skysqlcli_get_credentials.rst>`_ 	 - Retrieve default service credentials
* `skysqlcli get providers <skysqlcli_get_providers.rst>`_ 	 - Retrieve list of cloud providers
* `skysqlcli get quotas <skysqlcli_get_quotas.rst>`_ 	 - Retrieve quota information
* `skysqlcli get regions <skysqlcli_get_regions.rst>`_ 	 - Retrieve list of provider regions
* `skysqlcli get service-types <skysqlcli_get_service-types.rst>`_ 	 - Retrieve MariaDB SkySQL service-type information
* `skysqlcli get services <skysqlcli_get_services.rst>`_ 	 - Retrieve service information
* `skysqlcli get sizes <skysqlcli_get_sizes.rst>`_ 	 - Retrieve list of machine sizes
* `skysqlcli get status <skysqlcli_get_status.rst>`_ 	 - Retrieve current status for a service
* `skysqlcli get tiers <skysqlcli_get_tiers.rst>`_ 	 - Retrieve list of MariaDB SkySQL account tiers
* `skysqlcli get topologies <skysqlcli_get_topologies.rst>`_ 	 - Retrieve list of MariaDB SkySQL service topologies
* `skysqlcli get versions <skysqlcli_get_versions.rst>`_ 	 - Retrieve list of MariaDB SkySQL service versions

