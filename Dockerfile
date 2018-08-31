FROM golang:latest

RUN go get github.com/go-telegram-bot-api/telegram-bot-api

RUN mkdir /go/src/tg-bot-endless-story

ADD . /go/src/tg-bot-endless-story/
WORKDIR /go/src/tg-bot-endless-story

RUN go build -o main Main.go

VOLUME /go/src/tg-bot-endless-story/data

CMD ["/go/src/tg-bot-endless-story/main"]