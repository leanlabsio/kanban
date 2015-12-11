FROM scratch

EXPOSE 80

COPY bin/kanban_x86_64_linux /kanban

CMD ["/kanban"]
