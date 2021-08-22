# test_db [![Docker Automated build](https://img.shields.io/docker/automated/aleksi/test_db.svg)](https://hub.docker.com/r/aleksi/test_db/)

test_db is a repository and a Docker image with various example databases:

* MySQL
  * [World and World X (for MySQL X plugin)](https://dev.mysql.com/doc/world-setup/en/)
  * [Sakila](https://dev.mysql.com/doc/sakila/en/)
  * [datacharmer/test_db a.k.a. Employees](https://github.com/datacharmer/test_db)
  * [Menagerie](https://dev.mysql.com/doc/index-other.html)
* PostgreSQL
  * [World](http://pgfoundry.org/projects/dbsamples/)
  * [Pagila](https://github.com/devrimgunduz/pagila)

They can be used for example, tests, etc.

Docker image is build on top of busybox to make it explorable. It also exposes `/test_db` as a volume.


## Changelog

### v1.2.0 - not released yet

* Updated MySQL databases:
  * World and World X [from September 2016 to December 2019](https://dev.mysql.com/doc/world-setup/en/world-setup-history.html);
  * Sakila [from 1.0 to 1.2](https://dev.mysql.com/doc/sakila/en/sakila-news.html);
  * datacharmer/test_db a.k.a. Employees to [1.0.7](https://github.com/datacharmer/test_db/blob/master/Changelog);
* Updated PostgreSQL databases:
  * Pagila [from 0.10.1 to 2.1.0](https://github.com/devrimgunduz/pagila#version-history).

### v1.1.0 - 2019-06-13

* Added PostgreSQL databases: World and Pagila.

### v1.0.0 - 2018-02-21

* Initial version.
