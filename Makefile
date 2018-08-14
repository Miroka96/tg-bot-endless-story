NAME = endless-story
TAG = 1.0
IMAGE = $(NAME):$(TAG)
VOLUME = endless-story-data

.PHONY: build-container test-container deploy-container clean

build-container:
	docker build -t $(IMAGE) .

test-container:
	docker run -v $(VOLUME):/go/src/tg-bot-endless-story/data $(IMAGE)

deploy-container:
	docker run --detach --restart always --name=$(NAME) -v $(VOLUME):/go/src/tg-bot-endless-story/data $(IMAGE)

clean:
	docker volume rm $(VOLUME)
	docker rm $(NAME)