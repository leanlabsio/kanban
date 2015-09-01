FROM alpine:3.2

EXPOSE 80

ENV GITLAB_HOST=https://gitlab.com

WORKDIR /app
COPY ./kanban ./

ENTRYPOINT ["/bin/sh", "-c"]

CMD ["./kanban daemon --ip 0.0.0.0 --port 80 --gh $GITLAB_HOST"]
