FROM scratch

EXPOSE 80

COPY rel/kanban_x86_64_linux /kanban

CMD ["/kanban", "server"]
