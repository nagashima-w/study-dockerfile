FROM golang:1.11.4-alpine3.7
RUN apk update && \
    rm -rf /var/cache/apk/*
COPY src/main.go .
RUN go build main.go
EXPOSE 3000
CMD [ "./main" ]