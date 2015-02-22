### Installation

The easiest way to deploy Leanlabs Kanban board is to use [fig](http://www.fig.sh/). Assuming you have installed fig and Docker.

Change default environment variables defined in fig.yml then run 

```bash
fig up
```

### Configuration

Configuration of board provided through docker environment variables:

GITLAB_HOST - Your Gitlab installation host
GITLAB_API_TOKEN - Your Gitlab private API token, if defined used by default for all API requests
GITLAB_OAUTH_CLIENT_ID - Application ID
GITLAB_OAUTH_CLIENT_SECRET - Application secret
GITLAB_BASIC_LOGIN, GITLAB_BASIC_PASSWORD - HTTP basic authentication login and password, if you use it
KANBAN_SECRET_KEY - Token used to sign boards JSON Web Token

