FROM busybox

ADD mongodb /test_db/mongodb
ADD mysql /test_db/mysql
ADD postgresql /test_db/postgresql

VOLUME /test_db

LABEL org.opencontainers.image.source=https://github.com/AlekSi/test_db
LABEL org.opencontainers.image.title=test_db
