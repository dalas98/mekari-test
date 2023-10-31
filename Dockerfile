FROM golang:alpine as Golang

RUN apk update && apk add --no-cache git

COPY . .

RUN GOPATH= go test ./... -v

RUN GOPATH= go build -o /main main.go

FROM alpine:latest

COPY --from=Golang /go/app/repositories/db/migrations /app/repositories/db/migrations

COPY --from=Golang /main .

EXPOSE 8080

ENTRYPOINT ["./main"]