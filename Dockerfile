FROM golang:1.14.4-alpine3.12

COPY go-gmux-proper-unit-testing-api go-gmux-proper-unit-testing-api

ENTRYPOINT ["./go-gmux-proper-unit-testing-api"]