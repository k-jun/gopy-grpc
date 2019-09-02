FROM golang:1.12.7-alpine3.9 as builder
ENV GO111MODULE=on

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev make

WORKDIR /app
COPY . .
RUN go build -o /server ./main.go

FROM alpine:3.9
EXPOSE 8080
RUN apk update && apk upgrade && \
    apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /server ./app/server
# USER nobody

CMD ["./app/server"]