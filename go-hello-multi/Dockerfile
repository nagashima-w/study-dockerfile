FROM golang:1.11.4-alpine3.7 as builder
RUN apk update && \
    rm -rf /var/cache/apk/*
COPY src/main.go .
RUN go build main.go

FROM alpine:latest as runner
RUN apk update && \
    rm -rf /var/cache/apk/*
COPY --from=builder /go/main ./main
EXPOSE 3000
CMD [ "./main" ]