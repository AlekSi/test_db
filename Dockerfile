FROM busybox

ADD mysql /test_db/mysql
ADD postgresql /test_db/postgresql

VOLUME /test_db
