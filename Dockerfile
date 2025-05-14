
# Use the official Postgres 15 base image
FROM postgres:15

# Set default env vars (can still be overridden at runtime)
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

# Install PostGIS and its dependencies
RUN apt-get update \
    && apt-get install -y postgis postgresql-15-postgis-3 \
    && rm -rf /var/lib/apt/lists/*
    
# Optional: add SQL scripts to auto-run on first container startup
COPY ./seeds.sql /docker-entrypoint-initdb.d/


COPY ./pg_hba.conf /etc/postgresql/pg_hba.conf
COPY ./postgresql.conf /etc/postgresql/postgresql.conf


# Set Postgres to use the custom config files
CMD ["postgres", "-c", "config_file=/etc/postgresql/postgresql.conf", "-c", "hba_file=/etc/postgresql/pg_hba.conf"]
