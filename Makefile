types:
		~/go/bin/tygo generate 

front:
		cd frontend && yarn dev

back:
		PORT=1337 go run ./backend

deploy:
		cd backend && git push heroku

all: types front back
