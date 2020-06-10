build:
	docker build -f alpine.dockerfile -t install-it-alpine .

run: build
	docker run --rm -it install-it-alpine

build-ubuntu:
	docker build -f ubuntu.dockerfile -t install-it-ubuntu .

run-ubuntu: build-ubuntu
	docker run --rm -it install-it-ubuntu

