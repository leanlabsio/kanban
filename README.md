# GitLab issues made awesome

[![Join the chat at https://gitter.im/leanlabsio/kanban](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/leanlabsio/kanban?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Analytics](https://ga-beacon.appspot.com/UA-66361671-1/leanlabs/kanban)](https://github.com/igrigorik/ga-beacon)
#### Instant project management for your GitLab repositories

# FAQ

1. [How to install Kanban.Leanlabs](https://gitlab.com/leanlabsio/kanban/wikis/Installation)
2. [How to customize column](https://gitlab.com/leanlabsio/kanban/wikis/Customizing-columns)


## Installation

The easiest way to deploy Leanlabs Kanban board is to use docker-compose. 
Assuming you have installed [docker-compose](https://docs.docker.com/compose/) and [Docker](https://www.docker.com/).

### 1. Installation with Docker

> git clone https://gitlab.com/leanlabsio/kanban.git
>
> cd kanban

#### 1.1 Register GitLab Application for OAuth to work

Go to https://gitlab.com/profile/applications or your GitLab installation and register your application to get the application client ID and client secret key required for OAuth.

**Where**

> Redirect url http://{your-kanban-installation-host}/assets/html/user/views/oauth.html

#### 1.2 Change default environment variables defined in docker-compose.yml 

**Where**

> KANBAN_SERVER_HOSTNAME - URL on which LeanLabs Kanban will be reachable, required
>
> KANBAN_SECURITY_SECRET - This string is used to generate user auth tokens
>
> KANBAN_GITLAB_URL - Your GitLab host URL, required
>
> KANBAN_GITLAB_CLIENT - Your GitLab OAuth client ID, required for OAuth to work
>
> KANBAN_GITLAB_SECRET - Your GitLab OAuth client secret key, required for OAuth to work

**Then**

> docker-compose up -d


## Upgrading

If you followed instructions from "Installation with Docker", then the easiest way to upgrade would be:

> git pull
>
> docker-compose up -d

