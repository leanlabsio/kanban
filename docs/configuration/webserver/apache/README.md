# Apache web server configuration guide

The configuration files in this directory have been tested with Apache version 2.4.7.

To set up kanban you first need to enable several additional apache modules:

- proxy_wstunnel (to proxy web sockets)
- proxy
- proxy_http

## Common issues

Internally Kanban uses GitLab API. So misconfiguration of the GitLab web server could lead to potential kanban malfunctions.
Here you will find common issues we faced with GitLab + Apache setup and how to solve them.

1. Boards are not showing up.

    The symptoms - you can login to kanban, you can see the projects list in kanban, but can not open any board.

    First of all check your GitLab apache config contains rewrite rules like this one:

    ```
    RewriteRule .* http://127.0.0.1:8080%{REQUEST_URI} [P,QSA]
    ```

    Note the flags [P,QSA] - these are not enough for GitLab API to work properly, there must be NE flag present, so the proper rule is:

    ```
    RewriteRule .* http://127.0.0.1:8080%{REQUEST_URI} [NE,P,QSA]
    ```

    Also, please check the [gitlab recipes repo](https://github.com/gitlabhq/gitlab-recipes/tree/master/web-server/apache) for the latest apache configs.

2. Kanban reloads in a loop.

    The symptoms - you can login to kanban, but then the browser keep reloading the page.

    This could be caused by websockets proxy misconfiguration. Kanban listens for websocket connections on /ws/ location,
    so it should be proxied properly and there must be Host, Connection and Upgrade headers present.

3. Config example

    ```apache
    <VirtualHost *:80>
        ProxyPreserveHost On

        <Location />
                ProxyPass http://127.0.0.1:9000/
                ProxyPassReverse http://127.0.0.1:9000/
        </Location>

        <Location /ws/>
            ProxyPass ws://127.0.0.1:9000/ws/
            ProxyPassReverse ws://127.0.0.1:9000/ws/
        </Location>
    </VirtualHost>
    ```

    [view on github](https://github.com/leanlabsio/kanban/blob/master/docs/configuration/webserver/apache/apache.conf)
