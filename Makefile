IMAGE = leanlabs/client
TAG   = 1.3.0

help:
	@echo "Here will be brief doc"

test:
	@docker run -d -P --name selenium-hub selenium/hub:2.47.1
	@docker run -d --link selenium-hub:hub selenium/node-chrome:2.47.1
	@docker run -d --link selenium-hub:hub selenium/node-firefox:2.47.1
	@docker run -d -P $(IMAGE):$(TAG)
	@protractor $(CURDIR)/tests/e2e.conf.js

build:
	@docker run --rm -v $(CURDIR):/data -v $$HOME/node_cache:/cache leanlabs/npm-builder npm install
	@docker run --rm -v $(CURDIR):/data -v $$HOME/node_cache:/cache leanlabs/npm-builder bower install --allow-root
	@docker run --rm -v $(CURDIR):/data -v $$HOME/node_cache:/cache leanlabs/npm-builder grunt build
	@go-bindata -pkg=tempates -o templates/templates.go templates/...
	@go-bindata -pkg=web -o web/web.go web/...
	@GOOS=linus CGO_ENABLED=0 go build -a -o kanban

release:
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)
	@docker push $(IMAGE):latest
	@docker push $(IMAGE):$(TAG)

.PHONY: help test build release
