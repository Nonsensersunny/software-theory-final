FROM golang:latest AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod vendor && go build cmd/main.go

FROM debian:latest
COPY --from=build /app/main .
CMD ["./main"]