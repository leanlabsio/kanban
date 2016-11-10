all: build

help:
	@echo "build - create release for production with compiled docs\n" \
	      "start - start development environment"

build:
	@rm -rf docs/*

	@docker run --rm \
		-v $(CURDIR):/data \
		-w /data \
		leanlabs/git submodule init

	@docker run --rm \
		-v $(CURDIR):/data \
		-w /data \
		leanlabs/git submodule update

	@docker run --rm \
		-w /data/kanban.docs \
		-v $(CURDIR):/data \
		leanlabs/git pull origin master

	@docker run --rm \
		-v $(CURDIR):/data cnam/md2html \
		-s build/_Sidebar.md \
		-p /docs \
		-o docs \
		-t build/templates/documentation.tpl \
		-i kanban.docs/docs

start:
	@docker-compose up -d

stop:
	@docker-compose stop

.PHONY: build
