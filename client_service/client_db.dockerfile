FROM postgres:15.6

ADD ./client_service/sql/create.sql /docker-entrypoint-initdb.d/
