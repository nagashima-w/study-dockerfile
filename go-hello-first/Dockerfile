FROM golang:1.11.4
COPY src/main.go .
RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get autoremove -y && \
    apt-get autoclean -y && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*
RUN go build main.go
EXPOSE 3000
CMD [ "./main" ]