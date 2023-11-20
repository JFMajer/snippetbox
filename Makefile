# Makefile for managing Docker Compose actions

# Start up the application and its dependencies
up:
	docker-compose up -d

# Shut down the application and its dependencies
down:
	docker-compose down

# Rebuild the application container
build:
	docker-compose build

# Display the status of the containers
status:
	docker-compose ps

# View application logs
logs:
	docker-compose logs -f snippetbox
