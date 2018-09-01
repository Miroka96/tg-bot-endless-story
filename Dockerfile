FROM golang:latest

RUN go get github.com/go-telegram-bot-api/telegram-bot-api

RUN mkdir /tg-bot-endless-story

ADD bot/ /tg-bot-endless-story/bot/
ADD conf/ /tg-bot-endless-story/conf/
ADD Main.go /tg-bot-endless-story/

WORKDIR /tg-bot-endless-story

RUN go build -o main .

VOLUME /tg-bot-endless-story/data

CMD ["/tg-bot-endless-story/main"]
