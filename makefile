.PHONY: build, test, run, migrate_up, migrate_down
run: 
	./url_translater

build:
	# go mod download && 
	go build -o url_translater ./cmd/main.go

test:
	go test -v -race -timeout 30s ./...

migrate_up:
	migrate -path ./migrations -database 'postgres://pqkcourpqdzses:4c456c6e0a0ea78a8681af1da6297d71d92f158dd9b5c5fcf851b83140d15509@ec2-54-228-174-49.eu-west-1.compute.amazonaws.com:5432/d2n8bu73v867ev' up
migrate_down:
	migrate -path ./migrations -database 'postgres://pqkcourpqdzses:4c456c6e0a0ea78a8681af1da6297d71d92f158dd9b5c5fcf851b83140d15509@ec2-54-228-174-49.eu-west-1.compute.amazonaws.com:5432/d2n8bu73v867ev' down

.DEFAULT_GOAL := build