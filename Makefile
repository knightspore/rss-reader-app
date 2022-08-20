types:
		~/go/bin/tygo generate 

fe:
		cd frontend && yarn dev

be:
		go run ./backend

all: types fe be
