FROM debian:jessie

WORKDIR /bin/app
ADD http-trash /bin/app/main

EXPOSE 8080

CMD [ "/bin/app/main" ]