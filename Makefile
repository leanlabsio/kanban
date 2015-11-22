IMAGE = leanlabs/kanban
TAG   = 1.4.2

help:
	@echo "Here will be brief doc"

test:
	@docker run -d -P --name selenium-hub selenium/hub:2.47.1
	@docker run -d --link selenium-hub:hub selenium/node-chrome:2.47.1
	@docker run -d --link selenium-hub:hub selenium/node-firefox:2.47.1
	@docker run -d -P $(IMAGE):$(TAG)
	@protractor $(CURDIR)/tests/e2e.conf.js

build:
	@docker run --rm \
		-v $(CURDIR):/data \
		-v $$HOME/node_cache:/cache \
		-v /etc/passwd:/etc/passwd \
		-v /etc/group:/etc/group \
		-u $$USER \
		leanlabs/npm-builder npm install

	@docker run --rm \
		-v $(CURDIR):/data \
		-v $$HOME/node_cache:/cache \
		-v /etc/passwd:/etc/passwd \
		-v /etc/group:/etc/group \
		-e HOME=/cache \
		-u $$USER \
		leanlabs/npm-builder bower install

	@docker run --rm \
		-v $(CURDIR):/data \
		-v $$HOME/node_cache:/cache \
		-v /etc/passwd:/etc/passwd \
		-v /etc/group:/etc/group \
		-u $$USER \
		leanlabs/npm-builder grunt build

templates/templates.go: $(find $(CURDIR)/templates -name "*.html" -type f)
	@go-bindata -pkg=templates \
		-o templates/templates.go \
		templates/...

web/web.go: $(find $(CURDIR)/web/ -name "*" ! -name "web.go" -type f)
	@go-bindata -pkg=web \
		-o web/web.go \
		web/assets/... web/images/... web/template/...

kanban: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		leanlabs/golang-builder


bin/linux/x86_64/kanban: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		leanlabs/golang-builder

	-mkdir -p $(CURDIR)/bin/linux/x86_64/
	@mv kanban $(CURDIR)/bin/linux/x86_64/

bin/darwin/x86_64/kanban: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		-e GOOS=darwin \
		-e GOARCH=amd64 \
		leanlabs/golang-builder

	-mkdir -p $(CURDIR)/bin/darwin/x86_64/
	@mv kanban $(CURDIR)/bin/darwin/x86_64/

release: clean build templates/templates.go web/web.go kanban
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)
	@docker push $(IMAGE):latest
	@docker push $(IMAGE):$(TAG)

clean:
	@rm -rf web/
	@rm -f templates/templates.go
	@rm -f kanban

.PHONY: help test build release
