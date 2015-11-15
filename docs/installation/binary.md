# Installation from binary

### 1. Download

> [Linux x86_64](https://github.com/leanlabsio/kanban/releases/download/1.4.1/kanban_x86_64_linux)
>
> [MacOS x86_64](https://github.com/leanlabsio/kanban/releases/download/1.4.1/kanban_x86_64_darwin)

#### 1.2 Install redis

[Redis](http://redis.io/download#installation)

#### 1.3 Start server

> chmod +x kanban
>
> ./kanban server --redis-addr 127.0.0.1:6379

#### 1.4 Command line flags

List available cli flags:

> ./kanban server -h
