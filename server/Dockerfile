FROM registry.semaphoreci.com/golang:1.18 as builder

ENV APP_HOME /go/src/chatserver

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o chatserver ./cmd/chatserver/main.go

FROM registry.semaphoreci.com/golang:1.18

ENV APP_HOME /go/src/chatserver
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/chatserver $APP_HOME

EXPOSE 8080

CMD ["./chatserver"]