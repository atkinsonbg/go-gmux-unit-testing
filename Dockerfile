FROM scratch

COPY go-gmux-proper-unit-testing-api go-gmux-proper-unit-testing-api

ENTRYPOINT ["./go-gmux-proper-unit-testing-api"]