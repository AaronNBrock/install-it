build:
	docker build -t install-it .

run: build
	docker run --rm -it install-it