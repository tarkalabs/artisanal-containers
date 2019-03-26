.PHONY: clean
.PHONY: build
default: clean docker-build

CREATE_CONTAINER=$(shell docker create artisanal-containers)
SET_CONTAINER_ID=$(eval CONTAINER_ID=$(CREATE_CONTAINER))

clean: 
	rm -rf build/artisanal-containers

docker-build:
	docker build -t artisanal-containers .
	$(eval CONTID=$(docker create artisanal-containers))
	$(SET_CONTAINER_ID)
	docker cp $(CONTAINER_ID):go/src/github.com/tarkalabs/artisanal-containers/artisanal-containers build/artisanal-containers
	docker rm $(CONTAINER_ID)