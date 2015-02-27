
# GitLab issues made of awesome

#### Instant project management for your GitLab repositories

# FAQ

1. [How to customize column](customise_columns)
2. [How install Kanban.Leanlabs](install)

## Installation Kanban.LeanLabs

The easiest way to deploy Leanlabs Kanban board is to use fig. 
Assuming you have installed [fig](http://www.fig.sh/) and [Docker](https://www.docker.com/).

### 1. Symple usage

> git clone https://github.com/leanlabsio/kanban.git
>
> cd kanban

Change default environment variables defined in fig.yml 

**Where**

Main variables

> GITLAB_HOST - Your Gitlab installation host, required
>
> KANBAN_SECRET_KEY - Your Random secret key usage for generate jwt token, required
>
> GITLAB_API_TOKEN - Your Gitlab private API token, Using for gitlab api request for all users

**Then**

> fig up -d

### 2. Register App For GitLab OAuth

Go to https://gitlab.com/profile/applications or you installation gitlab and register your application to get the application keys needed for OAuth.

**Where**

> Redirect url https://{your-host}/assets/html/user/views/oauth.html

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

If you gitlab instance usage basic authentication set variables 

> GITLAB_BASIC_LOGIN - HTTP basic authentication login
>
> GITLAB_BASIC_PASSWORD -  HTTP basic authentication password


> ## If your usage in Kanban.LeanLabs Basic authentication OAuth not working.
