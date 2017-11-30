FROM golang

ADD . /go/src/github.com/dzyanis/go-service-skeleton

RUN go get github.com/gin-gonic/gin && go install github.com/dzyanis/go-service-skeleton

ENV SERVICE_PORT 3000

ENTRYPOINT /go/bin/go-service-skeleton

EXPOSE 3000