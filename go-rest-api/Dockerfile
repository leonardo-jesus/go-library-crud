# Dockerfile.development
FROM golang:1.19 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /api

CMD ["air"]

# Dockerfile.production
FROM golang:1.19-alpine as built

WORKDIR /api
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /usr/bin/main ./*.go

FROM busybox

EXPOSE 3333
COPY --from=built /usr/bin/main /usr/bin/main
CMD ["main"]