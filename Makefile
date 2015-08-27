help:
	@echo "build - create release for production with compiled docs\n" \
		  "start - start development environment"

build:
	@rm -rf documentation/*
	@docker run --rm -v $(CURDIR):/data leanlabs/git-builder submodule init
	@docker run --rm -v $(CURDIR):/data leanlabs/git-builder submodule update
	@docker run --rm -w /data/kanban.wiki -v $(CURDIR):/data leanlabs/git-builder pull origin master
	@docker run --rm -w /data -v $(CURDIR):/data alpine ./build/md2html

start:
	@docker-compose up -d

stop:
	@docker-compose stop

.PHONY: build
