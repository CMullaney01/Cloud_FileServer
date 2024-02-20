@echo off
docker compose run --rm start_dependencies
docker compose up -d mongodb
docker compose up keycloak