FROM golang:latest AS build
RUN mkdir -p /go/src/software-theory-final
ADD . /go/src/software-theory-final
WORKDIR /go/src/software-theory-final
ENV GOPROXY=https://goproxy.io
RUN go mod vendor && go build cmd/main.go

FROM debian:latest
RUN mkdir /app
WORKDIR /app
COPY --from=build /go/src/software-theory-final/main .
COPY --from=build /go/src/software-theory-final/config/ .
CMD ["./main"]