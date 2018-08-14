FROM golang:latest

RUN mkdir /go/src/tg-bot-endless-story

ADD . /go/src/tg-bot-endless-story/
WORKDIR /go/src/tg-bot-endless-story

RUN go get github.com/go-telegram-bot-api/telegram-bot-api

RUN go build -o main .

VOLUME /go/src/tg-bot-endless-story/data

CMD ["/go/src/tg-bot-endless-story/main"]