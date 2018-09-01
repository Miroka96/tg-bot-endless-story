NAME = endless-story
TAG = 1.0
IMAGE = $(NAME):$(TAG)
VOLUME = endless-story-data
MOUNTPATH = /tg-bot-endless-story/data

.PHONY: build-container test-container build-test-container deploy-container undeploy-container redeploy-container build-redeploy-container clean install-dependencies configure

build-container:
	docker build -t $(IMAGE) .

test-container:
	docker run -v $(VOLUME):$(MOUNTPATH) $(IMAGE)

build-test-container: build-container test-container

deploy-container:
	docker run --detach --restart always --name=$(NAME) -v $(VOLUME):$(MOUNTPATH) $(IMAGE)

undeploy-container:
	-docker stop $(NAME)
	docker rm $(NAME)

redeploy-container: undeploy-container deploy-container

build-redeploy-container: build-container redeploy-container

clean:
	-docker volume rm $(VOLUME)
	-docker rm $(NAME)

install-dependencies:
	go get -u github.com/go-sql-driver/mysql
	go get github.com/go-telegram-bot-api/telegram-bot-api

configure:
	sh configure.sh
