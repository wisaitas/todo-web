FROM postgres:17

COPY deployment/docker-image/scripts /docker-entrypoint-initdb.d/

ENV POSTGRES_PASSWORD=root

EXPOSE 5432
