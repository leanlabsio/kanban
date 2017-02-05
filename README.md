# GitLab issues made awesome

[![Join the chat at https://gitter.im/leanlabsio/kanban](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/leanlabsio/kanban?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Analytics](https://ga-beacon.appspot.com/UA-66361671-1/leanlabs/kanban)](https://github.com/igrigorik/ga-beacon)
[![](https://badge.imagelayers.io/leanlabs/kanban:1.4.0.svg)](https://imagelayers.io/?images=leanlabs/kanban:1.4.0 'Get your own badge on imagelayers.io')
#### Instant project management for your GitLab repositories

## Installation

Minimum Install Requrements:  
OS: kernel minimum 3.10 (centOS 7, Ubuntu 14.04)  
Packages: git, curl  

>`sudo yum -y install git, curl`  

The easiest way to deploy Leanlabs Kanban board is to use docker-compose. Install instructions here.
Assuming you have installed [Docker](http://docs.docker.com/engine/installation/) and [docker-compose](http://docs.docker.com/compose/install/).

### 1. Installation with Docker

>` git clone https://gitlab.com/leanlabsio/kanban.git`
>
>` cd kanban`


#### 1.1 Register GitLab Application for OAuth to work

Go to https://gitlab.com/profile/applications or your GitLab installation and register your application to get the application client ID and client secret key required for OAuth.

**Where**

> `Redirect url http[s]://{KANBAN_SERVER_HOSTNAME}/assets/html/user/views/oauth.html`

#### 1.2 Change default environment variables defined in docker-compose.yml 

**Where**

> `KANBAN_SERVER_HOSTNAME` | http[s]://{KANBAN_SERVER_HOSTNAME} - URL on which LeanLabs Kanban will be reachable [same as redirect url with out /assets/html...], required
>
> `KANBAN_SECURITY_SECRET` | Change this string to antyhing you like. This string is used to generate user auth tokens
>
> `KANBAN_GITLAB_URL` | http[s]://{gitlab.example.com:port} - Your GitLab host URL, required
>
> `KANBAN_GITLAB_CLIENT` | Your GitLab OAuth client application ID, required for OAuth to work. Git this from your gitlab server.
>
> `KANBAN_GITLAB_SECRET` | Your GitLab OAuth client secret key, required for OAuth to work. Git this from your gitlab server.
>
> `KANBAN_ENABLE_SIGNUP` | Wheter to enable sign up with user API token.

**Then**

> `docker-compose up -d`


## Upgrading

If you followed instructions from "Installation with Docker", then the easiest way to upgrade would be:

> `git pull`
>
> `docker-compose up -d`

## Changelog

You can view the changelog [here](https://gitlab.com/leanlabsio/kanban/blob/master/CHANGELOG.md)

# FAQ

1. [How to install Kanban.Leanlabs](http://kanban.leanlabs.io/docs/installation/)
2. [How to customize column](http://kanban.leanlabs.io/docs/usage/customize-columns)
