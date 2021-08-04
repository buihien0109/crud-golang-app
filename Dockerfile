# build stage
FROM golang:1.15-alpine AS build-env

ENV WDIR crud-app

WORKDIR /$WDIR

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' 

# final stage
FROM alpine:latest

RUN mkdir -p app/ /app/static /app/error

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build-env /crud-app/crud-app /app/

ENTRYPOINT ["./crud-app"]

EXPOSE 3000
