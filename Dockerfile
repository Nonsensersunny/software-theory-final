FROM golang:latest AS build
RUN mkdir -p /go/src/software-theory-final
ADD . /go/src/software-theory-final
WORKDIR /go/src/software-theory-final
ENV GOPROXY=https://goproxy.io
RUN go mod vendor && go build cmd/main.go

FROM python:latest
RUN mkdir -p /app/algorithm
WORKDIR /app
COPY --from=build /go/src/software-theory-final/main .
COPY --from=build /go/src/software-theory-final/config/ .
COPY --from=build /go/src/software-theory-final/algorithm/ algorithm/
RUN cd algorithm && \
    pip install --upgrade pip && \
    pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
CMD ["/app/main"]