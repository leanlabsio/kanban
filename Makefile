IMAGE = leanlabs/kanban
TAG   = 1.4.3

all: clean

test:
	@docker run -d -P --name selenium-hub selenium/hub:2.47.1
	@docker run -d --link selenium-hub:hub selenium/node-chrome:2.47.1
	@docker run -d --link selenium-hub:hub selenium/node-firefox:2.47.1
	@docker run -d -P $(IMAGE):$(TAG)
	@protractor $(CURDIR)/tests/e2e.conf.js

node_modules/: package.json
	@docker run --rm \
		-v $(CURDIR):/data \
		-v $$HOME/node_cache:/cache \
		-e HOME=/cache \
		leanlabs/npm-builder:latest npm install

bower_components/: bower.json
	@docker run --rm \
		-v $(CURDIR):/data \
		-v $$HOME/node_cache:/cache \
		-e HOME=cache \
		leanlabs/npm-builder bower install --allow-root

build: node_modules/ bower_components/
	@docker run --rm \
		-v $(CURDIR):/data \
		-v $$HOME/node_cache:/cache \
		-e HOME=/cache \
		leanlabs/npm-builder grunt build

templates/templates.go: $(shell find $(CURDIR)/templates -name "*.html" -type f)
	@docker run --rm \
		-v $(CURDIR):/data \
		leanlabs/go-bindata-builder \
		-pkg=templates -o templates/templates.go \
		templates/...

web/web.go: $(shell find $(CURDIR)/web/ -name "*" ! -name "web.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/data \
		leanlabs/go-bindata-builder \
		-pkg=web -o web/web.go \
		web/assets/... web/images/... web/template/...

rel/kanban_x86_64_linux: clean build templates/templates.go web/web.go $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/go/src/gitlab.com/leanlabsio/kanban \
		-w /go/src/gitlab.com/leanlabsio/kanban \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e GO15VENDOREXPERIMENT=1 \
		-e CGO_ENABLED=0 \
		--entrypoint=/usr/local/go/bin/go \
		leanlabs/golang-builder build -ldflags '-s' -v -o $@

rel/kanban_x86_64_darwin: clean build templates/templates.go web/web.go $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/go/src/gitlab.com/leanlabsio/kanban \
		-w /go/src/gitlab.com/leanlabsio/kanban \
		-e GOOS=darwin \
		-e GOARCH=amd64 \
		-e GO15VENDOREXPERIMENT=1 \
		-e CGO_ENABLED=0 \
		--entrypoint=/usr/local/go/bin/go \
		leanlabs/golang-builder build -ldflags '-s' -v -o $@

release: rel/kanban_x86_64_linux
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)
	@docker push $(IMAGE):latest
	@docker push $(IMAGE):$(TAG)

clean:
	@rm -rf web/
	@rm -f templates/templates.go
	@rm -f kanban

.PHONY: help test build release