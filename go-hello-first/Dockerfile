FROM golang:1.11.4
RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get autoremove -y && \
    apt-get autoclean -y && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*
COPY src/main.go .
RUN go build main.go
EXPOSE 3000
CMD [ "./main" ]