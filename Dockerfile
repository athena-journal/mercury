FROM alpine:latest

EXPOSE 8080

ADD mercury /bin/mercury

CMD ["mercury"]