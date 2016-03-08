IMAGE = leanlabs/kanban
TAG   = 1.5.0
CWD   = /go/src/gitlab.com/leanlabsio/kanban

all: clean

test:
	@docker run -d -P --name selenium-hub selenium/hub:2.47.1
	@docker run -d --link selenium-hub:hub selenium/node-chrome:2.47.1
	@docker run -d --link selenium-hub:hub selenium/node-firefox:2.47.1
	@docker run -d -P $(IMAGE):$(TAG)
	@protractor $(CURDIR)/tests/e2e.conf.js

node_modules/: package.json
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-v $$HOME/node_cache:/cache \
		-w $(CWD) \
		-e HOME=/cache \
		leanlabs/npm-builder:latest npm install

bower_components/: bower.json
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-v $$HOME/node_cache:/cache \
		-w $(CWD) \
		-e HOME=cache \
		leanlabs/npm-builder bower install --allow-root

build: node_modules/ bower_components/
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-v $$HOME/node_cache:/cache \
		-w $(CWD) \
		-e HOME=/cache \
		leanlabs/npm-builder grunt build

templates/templates.go: $(find $(CURDIR)/templates -name "*.html" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		leanlabs/go-bindata-builder \
		$(DEBUG) \
		-pkg=templates -o templates/templates.go \
		templates/...

web/web.go: $(find $(CURDIR)/web/ -name "*" ! -name "web.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		leanlabs/go-bindata-builder \
		$(DEBUG) \
		-pkg=web -o web/web.go \
		web/assets/... web/images/... web/template/...

rel/kanban_x86_64_linux: clean build templates/templates.go web/web.go $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e GO15VENDOREXPERIMENT=1 \
		-e CGO_ENABLED=0 \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.5.3 build -ldflags '-s' -ldflags "-X main.AppVer=$(TAG)" -v -o $@

rel/kanban_x86_64_darwin: clean build templates/templates.go web/web.go $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOOS=darwin \
		-e GOARCH=amd64 \
		-e GO15VENDOREXPERIMENT=1 \
		-e CGO_ENABLED=0 \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.5.3 build -ldflags '-s' -ldflags "-X main.AppVer=$(TAG)" -v -o $@

release: rel/kanban_x86_64_linux
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)
	@docker push $(IMAGE):latest
	@docker push $(IMAGE):$(TAG)

clean:
	@rm -rf web/
	@rm -f templates/templates.go

# Development targets
dev_redis:
	@docker inspect -f {{.State.Running}} kb_dev_redis || \
		docker run -d -p 6379:6379 --name kb_dev_redis leanlabs/redis

watch: build
	@docker inspect -f {{.State.Running}} kb_dev_watcher || \
		docker run -d \
			--name kb_dev_watcher \
			-v $(CURDIR):$(CWD) \
			-v $$HOME/node_cache:/cache \
			-w $(CWD) \
			-e HOME=/cache \
			leanlabs/npm-builder grunt watch

tmp/go/pkg/:
	@docker run --rm \
		-v $(CURDIR)/tmp/go/pkg:/go/pkg \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GO15VENDOREXPERIMENT=1 \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.5.3 install -v

dev : DEBUG=-debug

dev: watch templates/templates.go web/web.go dev_redis tmp/go/pkg/
	-docker rm -f kb_dev
	@docker run -d --link kb_dev_redis:redis --name kb_dev \
		-p 9000:80 \
		-v $(CURDIR):$(CWD) \
		-v $(CURDIR)/tmp/go/pkg:/go/pkg \
		-w $(CWD) \
		-e GO15VENDOREXPERIMENT=1 \
		-e KANBAN_ENABLE_SIGNUP=$(KANBAN_ENABLE_SIGNUP) \
		-e KANBAN_SERVER_HOSTNAME=$(KANBAN_SERVER_HOSTNAME) \
		-e KANBAN_GITLAB_CLIENT=$(KANBAN_GITLAB_CLIENT) \
		-e KANBAN_GITLAB_SECRET=$(KANBAN_GITLAB_SECRET) \
		-e KANBAN_REDIS_ADDR=redis:6379 \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.5.3 run -v main.go server

.PHONY: help test build release
