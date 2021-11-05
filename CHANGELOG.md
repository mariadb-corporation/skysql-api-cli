# v0.0.16 (Fri Nov 05 2021)

#### 🐛 Bug Fix

- Fix auto version included in go build of cli [#27](https://github.com/mariadb-corporation/skysql-api-cli/pull/27) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### 🚢 DevOps

- Determine version info at start of pipeline [#29](https://github.com/mariadb-corporation/skysql-api-cli/pull/29) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))
- Install additional plugins for `intuit/auto` in build pipeline [#28](https://github.com/mariadb-corporation/skysql-api-cli/pull/28) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.15 (Thu Nov 04 2021)

#### 🐛 Bug Fix

- Use kebab case for api-key flag in config file [#26](https://github.com/mariadb-corporation/skysql-api-cli/pull/26) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.14 (Wed Nov 03 2021)

#### 🚀 Enhancement

- Add `ssl_tls` field to create service requests [#25](https://github.com/mariadb-corporation/skysql-api-cli/pull/25) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.13 (Wed Oct 27 2021)

#### 🐛 Bug Fix

- Rename msvcid flag back to mdbid [#24](https://github.com/mariadb-corporation/skysql-api-cli/pull/24) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.12 (Mon Oct 25 2021)

#### 🚀 Enhancement

- Rename "database" to "service" for all resources [#23](https://github.com/mariadb-corporation/skysql-api-cli/pull/23) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.11 (Wed Oct 20 2021)

#### 📝 Documentation

- Add section to readme about install of the cli [#22](https://github.com/mariadb-corporation/skysql-api-cli/pull/22) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.10 (Tue Oct 12 2021)

#### 🚀 Enhancement

- Add tier flag to create database and get sizes commands [#21](https://github.com/mariadb-corporation/skysql-api-cli/pull/21) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

#### Authors: 1

- John Halbert ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

---

# v0.0.9 (Fri Oct 08 2021)

#### 🚀 Enhancement

- Add get status, create configuration, create allowed-address, update configuration, update status, delete configuration and delete allowed-address commands [#20](https://github.com/mariadb-corporation/skysql-api-cli/pull/20) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

#### Authors: 1

- John Halbert ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

---

# v0.0.8 (Fri Sep 24 2021)

#### 🐛 Bug Fix

- Use strings rather than pointers for newly required params [#19](https://github.com/mariadb-corporation/skysql-api-cli/pull/19) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))
- Avoid viper bug regarding flag name collisions [#18](https://github.com/mariadb-corporation/skysql-api-cli/pull/18) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.7 (Fri Sep 24 2021)

#### 🚀 Enhancement

- Add get versions command [#16](https://github.com/mariadb-corporation/skysql-api-cli/pull/16) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))
- Add get topologies command [#15](https://github.com/mariadb-corporation/skysql-api-cli/pull/15) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))
- Add get tiers command [#14](https://github.com/mariadb-corporation/skysql-api-cli/pull/14) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))
- Add get sizes command [#13](https://github.com/mariadb-corporation/skysql-api-cli/pull/13) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

#### Authors: 1

- John Halbert ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

---

# v0.0.6 (Fri Sep 24 2021)

#### 🚀 Enhancement

- Add get regions command [#12](https://github.com/mariadb-corporation/skysql-api-cli/pull/12) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))
- Add get providers command [#11](https://github.com/mariadb-corporation/skysql-api-cli/pull/11) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))
- Add get products command [#10](https://github.com/mariadb-corporation/skysql-api-cli/pull/10) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))
- Add get configurations command [#9](https://github.com/mariadb-corporation/skysql-api-cli/pull/9) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

#### Authors: 1

- John Halbert ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

---

# v0.0.5 (Fri Sep 24 2021)

#### 🚀 Enhancement

- Add get allowlist command [#8](https://github.com/mariadb-corporation/skysql-api-cli/pull/8) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

#### Authors: 1

- John Halbert ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

---

# v0.0.4 (Thu Sep 23 2021)

#### 🚀 Enhancement

- Update sdk to include `volume_iops` and `volume_type` for db [#17](https://github.com/mariadb-corporation/skysql-api-cli/pull/17) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.3 (Wed Sep 22 2021)

#### 🚀 Enhancement

- DBAAS-6841: Add create/update/delete database commands [#7](https://github.com/mariadb-corporation/skysql-api-cli/pull/7) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 1

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

---

# v0.0.2 (Tue Sep 21 2021)

#### 🚀 Enhancement

- DBAAS-7920: GET Billing [#6](https://github.com/mariadb-corporation/skysql-api-cli/pull/6) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

#### Authors: 1

- John Halbert ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

---

# v0.0.1 (Tue Sep 21 2021)

#### 💥 Breaking Change

- Introduce Constants [#5](https://github.com/mariadb-corporation/skysql-api-cli/pull/5) ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))

#### 🚀 Enhancement

- Introduce `skysqlcli get databases` as first command [#3](https://github.com/mariadb-corporation/skysql-api-cli/pull/3) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### 🚢 DevOps

- Introduce auto releases for windows/macos/linux [#4](https://github.com/mariadb-corporation/skysql-api-cli/pull/4) ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### ⚠️ Pushed to `main`

- Initial commit ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))

#### Authors: 2

- Don Mayo ([@mariadb-DonMayo](https://github.com/mariadb-DonMayo))
- John Halbert ([@mariadb-JohnHalbert](https://github.com/mariadb-JohnHalbert))
