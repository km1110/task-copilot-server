# develop
FROM golang:1.22 AS develop

WORKDIR /go/src/github.com/km1110/task-copilot

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]

## Build
FROM golang:1.22 AS build

WORKDIR /go/src/github.com/km1110/task-copilot

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /server ./cmd/api

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]
