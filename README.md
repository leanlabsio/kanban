
# GitLab issues made awesome

#### Instant project management for your GitLab repositories

# FAQ

1. [How to install Kanban.Leanlabs](https://gitlab.com/leanlabsio/kanban/wikis/install)
2. [How to customize column](https://gitlab.com/leanlabsio/kanban/wikis/Customizing-columns)


## Installation

The easiest way to deploy Leanlabs Kanban board is to use fig. 
Assuming you have installed [fig](http://www.fig.sh/) and [Docker](https://www.docker.com/).

### 1. Simple usage

> git clone https://gitlab.com/leanlabsio/kanban.git
>
> cd kanban

Change default environment variables defined in fig.yml 

**Where**

Main variables

> GITLAB_HOST - Your Gitlab installation host, required
>
> KANBAN_SECRET_KEY - Your Random secret key, used to generate jwt token, required
>
> GITLAB_API_TOKEN - Your Gitlab private API token, Using for gitlab api request for all users

**Then**

> fig up -d

### 2. Register App For GitLab OAuth

Go to https://gitlab.com/profile/applications or you installation gitlab and register your application to get the application keys needed for OAuth.

**Where**

> Redirect url http://{your-kanban-installation-host}/assets/html/user/views/oauth.html

### 3. Configure OAuth Environment

Change default environment variables defined in fig.yml 

> GITLAB_OAUTH_CLIENT_ID - Application ID
> 
> GITLAB_OAUTH_CLIENT_SECRET - Application Secret 

### 4. Upgrading Kanban.leanlabs

For upgrading Kanban LeanLabs to last version

> fig pull
> fig up -d

### 5. Basic Auth

If you gitlab installation secured with basic authentication

> GITLAB_BASIC_LOGIN - HTTP basic authentication login
>
> GITLAB_BASIC_PASSWORD -  HTTP basic authentication password

