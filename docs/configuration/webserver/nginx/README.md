# Nginx web server configuration guide

The configuration files in this directory have been tested with nginx 1.8+.

### Config example

```nginx
server {
    listen 80;

    location / {
        proxy_pass http://127.0.0.1:9000;
    }

    location /ws {
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Origin "";
        proxy_pass http://127.0.0.1:9000;
    }
}

```

[view on github](https://github.com/leanlabsio/kanban/blob/master/docs/configuration/webserver/nginx/nginx.conf)
