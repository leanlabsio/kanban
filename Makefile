help:
	@echo "build - create release for production with compiled docs \
 	      start - start development environment"

build:
	@docker run --it --rm 
start:
	@docker-compose up -d

stop:
	@docker-compose stop

