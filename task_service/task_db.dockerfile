FROM postgres:15.6

ADD ./task_service/sql/create.sql /docker-entrypoint-initdb.d/
