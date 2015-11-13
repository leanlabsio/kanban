### Apache web server configuration guide

The configuration files in this directory was tested with Apache version 2.4.7.

To set up kanban you first need to enable several additional apache modules:

- proxy_wstunnel (to proxy web sockets)
- proxy
- proxy_http

### Common issues

Internally Kanban uses GitLab API. So misconfiguration of GitLab web server could lead to potentially kanban malfunction.
Here you will find common issues we faced with GitLab + Apache setup and how to solve them.

1. Boards are not showing up.

    The simptoms - you could login to kanban, you could see projects list in kanban, but could not open any board.

    First of all check you GitLab apache config contains rewrite rules like this one:

    ```
    RewriteRule .* http://127.0.0.1:8080%{REQUEST_URI} [P,QSA]
    ```

    Note the flags [P,QSA] - these are not enough for GitLab API to work properly, there must be NE flag present, so the proper rule is:

    ```
    RewriteRule .* http://127.0.0.1:8080%{REQUEST_URI} [NE,P,QSA]
    ```

    Also, please check the [gitlab recipes repo](https://github.com/gitlabhq/gitlab-recipes/tree/master/web-server/apache) for latest apache configs.

2. Kanban reloads in a loop.

    The simptoms - you could login to kanban, but then browser keep reloading the page.

    This could be caused by websockets proxy misconfiguration. Kanban listens for websocket connections on /ws/ location,
    so it should be proxied properly and there must be Host, Connection and Upgrade headers present.
