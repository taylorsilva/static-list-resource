
test:
	go test -v -count=1 ./...

build:
	docker build -t taylorsilva/static-list-resource .
