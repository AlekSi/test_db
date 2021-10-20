# test_db

test_db is a repository and a [Docker image](https://hub.docker.com/r/aleksi/test_db/) with various example databases:

* MySQL
  * [World and World X (for MySQL X plugin)](https://dev.mysql.com/doc/world-setup/en/)
  * [Sakila](https://dev.mysql.com/doc/sakila/en/)
  * [datacharmer/test_db a.k.a. Employees](https://github.com/datacharmer/test_db)
  * [Menagerie](https://dev.mysql.com/doc/index-other.html)
* PostgreSQL
  * [World](http://pgfoundry.org/projects/dbsamples/)
  * [Pagila](https://github.com/devrimgunduz/pagila)
* MongoDB
  * Monila (converted from Pagila)

They can be used for example, tests, etc.

Docker image is build on top of busybox to make it explorable.
It also exposes `/test_db` as a volume that can be mounted into MySQL/PostgreSQL/MongoDB containers
for initializing databases on start-up. See [docker-compose.yml](docker-compose.yml) for examples and
images documentation for `docker-entrypoint-initdb.d` initialization scripts.


## Changelog

### v1.2.0 - 2021-10-20

* Updated MySQL databases:
  * World and World X [from September 2016 to December 2019](https://dev.mysql.com/doc/world-setup/en/world-setup-history.html);
  * Sakila [from 1.0 to 1.2](https://dev.mysql.com/doc/sakila/en/sakila-news.html):
    * files were renamed to `01-sakila-schema.sql` and `02-sakila-data.sql` to make them
      compatible with the official `mysql` Docker image's initialization scripts.
  * datacharmer/test_db a.k.a. Employees to [1.0.7](https://github.com/datacharmer/test_db/blob/master/Changelog);
* Updated PostgreSQL databases:
  * Pagila [from 0.10.1 to 2.1.0](https://github.com/devrimgunduz/pagila#version-history):
    * files were renamed to `01-pagila-schema.sql` and `02-pagila-data.sql` to make them
      compatible with the official `postgres` Docker image's initialization scripts;
    * `pagila-insert-data.sql` was removed as it contained the same data as (`02-`)`pagila-data.sql`;
    * please note that Sakila and Pagila databases have small differences, most notably in dates;
* Added MongoDB Monila database:
  * Converted from Pagila by a script in this repository.

### v1.1.0 - 2019-06-13

* Added PostgreSQL databases: World and Pagila.

### v1.0.0 - 2018-02-21

* Initial version.
