# Makefile - Appointment Platform

ENV_FILE=.env

## Run docker-compose up with build
up:
	docker-compose --env-file $(ENV_FILE) up --build

## Run docker-compose in detached mode
start:
	docker-compose --env-file $(ENV_FILE) up -d

## Stop all containers
stop:
	docker-compose down

## Show logs of all containers
logs:
	docker-compose logs -f

## Rebuild containers without cache
rebuild:
	docker-compose build --no-cache

## Run backend container with shell
backend-shell:
	docker exec -it backend sh

## Run frontend container with shell
frontend-shell:
	docker exec -it frontend sh

## Run cronjob container with shell
cronjob-shell:
	docker exec -it appointment-reminder-cron sh
