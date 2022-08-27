create.types:
		~/go/bin/tygo generate 

# Development

dev.frontend:
		cd frontend && yarn dev

dev.backend:
		PORT=1337 go run ./backend

# Build

build.backend:
	cd backend && go build -o ./../bin/backend_app

# Docker
docker:
	docker build -t readerbackend .

docker.dev:
	docker run -p 80:1337 -d readerbackend

# Testing
test.backend: 
		cd backend && go test . -v -cover
		cd backend/module && go test . -v -cover
		cd backend/parse && go test . -v -cover
		cd backend/util && go test . -v -cover
		cd backend/vo && go test . -v -cover

# Grouped Commands
test.all: test.backend

all: 
	make -j create.types dev.frontend dev.backend
