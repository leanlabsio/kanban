# Contributing

Contributions are **welcome** and will be fully **credited**.

We accept contributions via Pull Requests on [GitLab](https://gitlab.com/leanlabsio/kanban).

## Pull Requests

- **Document any change in behaviour** - Make sure the `README.md` and any other relevant documentation from directory /docs are kept up-to-date.

- **One pull request per feature** - If you want to do more than one thing, send multiple pull requests.

- **Send coherent history** - Make sure each individual commit in your pull request is meaningful. If you had to make multiple intermediate commits while developing, please [squash them](http://www.git-scm.com/book/en/v2/Git-Tools-Rewriting-History#Changing-Multiple-Commit-Messages) before submitting.

## Development

They are two ways for run the development environment - based on docker, or install required software on your development machine.

We are recommended use docker installation as we. We are test this installation on Mac os, Ubuntu, Centos, Debian.

### Running with docker

The following software is required to start for development environment:

- [Docker 1.6+](https://docker.io)
- GNU make

To start development environment you need:
    
1. Create gitlab application [How create application](http://kanban.leanlabs.io/docs/installation/docker#setting-up-oauth-via-gitlab)
2. Then run the following command in project root dir.

```bash
KANBAN_GITLAB_URL=https://gitlab.com KANBAN_GITLAB_CLIENT=client_id KANBAN_GITLAB_SECRET=secret_id make dev
```

That is it, now you have running development environment and could start writing code.

All dependenses will be gets automaticaly, after that kanban build go application and start listen on `127.0.0.1:9000` or IP you virtual machine with docker.

**NOTE** Golang server side application is not hot reloading, so any time you are changing backend code you should run the above command.


### Running without docker

We are doesn't recommended this process.

The following software is required for start development environment:

- [Node.js + npm](https://nodejs.org/en/)
- [Golang 1.6+](https://golang.org/)

To start development environment you need:

1. Create gitlab application [How create application](http://kanban.leanlabs.io/docs/installation/docker#setting-up-oauth-via-gitlab)
2. Install and start [redis server](http://redis.io) instance
3. The run the following commands in project root dir.

**NOTE** For using golang you need clone repository on **$GOPATH/src/gitlab.com/leanlabsio/kanban**


Install client dependencies

```bash
npm install

bower install --allow-root

grunt build
```

```bash
go get -u github.com/jteeuwen/go-bindata/...

go-bindata -debug -pkg=templates -o templates/templates.go templates/...

go-bindata -debug -pkg=web -o web/web.go web/...
```

``` bash
go run -v main.go server \
    --redis-addr REDIS_SERVER_IP:REDIS_SERVER_PORT \
    --kanban-gitlab-url https://gitlab.com\
    --kanban-gitlab-client client_id \
    --kanban-gitlab-secret secret_id
```

After that open New tab on terminal and execute

```
grunt watch

```

That's it, now you should have running dev env and could start writing code.
