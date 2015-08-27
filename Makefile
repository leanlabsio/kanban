help:
	@echo "build - create release for production with compiled docs\n" \
		  "start - start development environment"

build:
	@rm -rf documentation/*
	@docker run --rm -w /data/kanban.wiki -v $(CURDIR):/data leanlabs/git-builder pull origin master
	@docker run --rm -w /data -v $(CURDIR):/data alpine ./main

start:
	@docker-compose up -d

stop:
	@docker-compose stop

.PHONY: build
