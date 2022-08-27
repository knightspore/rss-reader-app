FROM golang:latest

RUN mkdir /build

RUN cd /build && git clone https://github.com/knightspore/rss-reader-app.git
RUN cd /build/rss-reader-app && make build.backend

WORKDIR /build/rss-reader-app/bin

EXPOSE 1337

CMD ["/build/rss-reader-app/bin/backend_app"]
