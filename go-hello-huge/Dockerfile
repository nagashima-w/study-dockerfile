FROM golang:1.11.4
COPY src/main.go .
RUN go build main.go
RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get autoremove -y
RUN apt-get autoclean -y
RUN rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*
EXPOSE 3000
CMD [ "./main" ]