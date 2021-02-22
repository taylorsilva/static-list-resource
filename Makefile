
test:
	go test ./...

build: test
	docker build -t taylorsilva/static-list-resource .
