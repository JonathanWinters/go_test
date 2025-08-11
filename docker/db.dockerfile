FROM postgres:13.4-alpine
COPY sql/init2.sql /sql/
COPY sql/schema.sql /sql/
WORKDIR /sql/
RUN cat init2.sql schema.sql > /docker-entrypoint-initdb.d/init.sql
EXPOSE 5432