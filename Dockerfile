FROM scratch

EXPOSE 80

COPY ./kanban ./

CMD ["./kanban"]
