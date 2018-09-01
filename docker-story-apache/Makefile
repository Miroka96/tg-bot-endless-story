NAME = endless-story-apache
TAG = 1.0
IMAGE = $(NAME):$(TAG)
VOLUME = endless-story-data
MOUNTPATH = /usr/local/apache2/htdocs
DATAMOUNT = $(VOLUME):$(MOUNTPATH):ro
HOSTPORT = 11000
PORTFORWARD = 127.0.0.1:$(HOSTPORT):80
DOCKERRUNOPTIONS = -v $(DATAMOUNT) -p $(PORTFORWARD)

.PHONY: build-container test-container build-test-container deploy-container build-deploy-container undeploy-container redeploy-container build-redeploy-container clean install-dependencies configure

build-container:
	docker build -t $(IMAGE) .

test-container:
	docker run $(DOCKERRUNOPTIONS) --rm $(IMAGE)

build-test-container: build-container test-container

deploy-container:
	docker run $(DOCKERRUNOPTIONS) --detach --restart always --name=$(NAME) $(IMAGE)

build-deploy-container: build-container deploy-container

undeploy-container:
	-docker stop $(NAME)
	docker rm $(NAME)

redeploy-container: undeploy-container deploy-container

build-redeploy-container: build-container redeploy-container

clean:
	-docker rm $(NAME)
