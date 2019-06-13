# test_db [![Docker Automated build](https://img.shields.io/docker/automated/aleksi/test_db.svg)](https://hub.docker.com/r/aleksi/test_db/)

test_db is a repository and a Docker image with various example databases:

* MySQL
  * [World and World X (for MySQL X plugin)](https://dev.mysql.com/doc/world-setup/en/)
  * [Sakila](https://dev.mysql.com/doc/sakila/en/)
  * [Employees](https://github.com/datacharmer/test_db)
  * [Menagerie](https://dev.mysql.com/doc/index-other.html)
* PostgreSQL
  * [World](http://pgfoundry.org/projects/dbsamples/)
  * [Pagila](http://pgfoundry.org/projects/dbsamples/)

They can be used for example, tests, etc.

Docker image is build on top of busybox to make it explorable. It also exposes `/test_db` as a volume.


## Changelog

### v1.1.0 - 2019-06-13

* Add PostgreSQL databases: World and Pagila.

### v1.0.0 - 2018-02-21

* Initial version.
