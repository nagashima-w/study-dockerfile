FROM golang:1.11.4-alpine3.7 as builder
COPY src/main.go .
RUN apk update && \
    rm -rf /var/cache/apk/*
RUN go build main.go

FROM alpine:latest as runner
COPY --from=builder /go/main ./main
RUN apk update && \
    rm -rf /var/cache/apk/*
EXPOSE 3000
CMD [ "./main" ]