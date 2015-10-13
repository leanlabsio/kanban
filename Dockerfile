FROM alpine:3.2

EXPOSE 80

ENV GITLAB_HOST=https://gitlab.com

WORKDIR /app
COPY ./kanban ./

ENTRYPOINT ["/bin/sh", "-c"]

CMD ["./kanban daemon --listen 0.0.0.0:80"]
