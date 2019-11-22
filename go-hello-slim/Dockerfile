FROM golang:1.11.4-alpine3.7
COPY src/main.go .
RUN apk update && \
    rm -rf /var/cache/apk/*
RUN go build main.go
EXPOSE 3000
CMD [ "./main" ]