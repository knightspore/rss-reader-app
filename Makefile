types:
		~/go/bin/tygo generate 

test: 
		cd backend && go test . -v -cover
		cd backend/module && go test . -v -cover
		cd backend/parse && go test . -v -cover
		cd backend/util && go test . -v -cover
		cd backend/vo && go test . -v -cover

front:
		cd frontend && yarn dev

back:
		PORT=1337 go run ./backend

deploy:
		cd backend && git push heroku

all: types front back
