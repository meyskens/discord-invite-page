FROM golang:1.15-alpine as build

RUN apk add --no-cache git

COPY ./ /go/src/github.com/meyskens/discord-join-page

WORKDIR /go/src/github.com/meyskens/discord-join-page

RUN go build -ldflags "-X main.revision=$(git rev-parse --short HEAD)" ./cmd/discord-join-page/

FROM alpine:3.12

RUN apk add --no-cache ca-certificates

RUN mkdir -p /go/src/github.com/meyskens/discord-join-page/
WORKDIR /go/src/github.com/meyskens/discord-join-page/
COPY ./www /go/src/github.com/meyskens/discord-join-page/www

COPY --from=build /go/src/github.com/meyskens/discord-join-page/discord-join-page /usr/local/bin/

CMD [ "/usr/local/bin/discord-join-page", "serve" ]
