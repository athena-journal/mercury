.PHONY: install build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

build: install
	go build -ldflags "-X main.version=$(TAG)" -o mercury .

serve: build
	./mercury

clean:
	rm ./mercury

pack:
	GOOS=linux make build
	docker build -t athena/mercury:$(TAG) .

deploy:
	envsubst < k8s/deployment.yaml | kubectl apply -f -

ship: pack deploy clean