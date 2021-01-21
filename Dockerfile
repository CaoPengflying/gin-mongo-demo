FROM golang:1.14-alpine as build

ENV GOPROXY=https://goproxy.io

WORKDIR /usr/src/app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download


COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o server main.go

FROM scratch as runner

COPY --from=build /usr/src/app/server /opt/app/

