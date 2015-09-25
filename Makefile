IMAGE = leanlabs/client
TAG   = 1.2.4

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
	@go-bindata -pkg=web -o web/web.go web/...

templates/templates.go: $(find $(CURDIR)/templates -name "*.html" -type f)
	@go-bindata -pkg=templates -o templates/templates.go templates/...

kanban: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm -v $(CURDIR):/data leanlabs/go-builder

release: build templates/templates.go kanban
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)
#	@docker push $(IMAGE):latest
# 	@docker push $(IMAGE):$(TAG)

clean:
	@rm -f web/web.go
	@rm -f templates/templates.go

.PHONY: help test build release
