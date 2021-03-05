
test:
	go test -count=1 ./...

build:
	docker build -t taylorsilva/static-list-resource .
