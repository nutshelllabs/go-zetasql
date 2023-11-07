docker/build:
	docker build --platform=linux/amd64 -t go-zetasql .

test/linux: docker/build
	docker run --rm go-zetasql bash -c "go test -v ."
