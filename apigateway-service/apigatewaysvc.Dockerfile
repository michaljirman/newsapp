FROM golang:1.13.4 as builder

COPY vendor /go/src/github.com/michaljirman/newsapp/vendor

COPY newsfeeder-service/pb /go/src/github.com/michaljirman/newsapp/newsfeeder-service/pb
COPY newsfeeder-service/pkg/ /go/src/github.com/michaljirman/newsapp/newsfeeder-service/pkg/

RUN mkdir -p /go/src/github.com/michaljirman/newsapp/apigateway-service
COPY apigateway-service /go/src/github.com/michaljirman/newsapp/apigateway-service
WORKDIR /go/src/github.com/michaljirman/newsapp/apigateway-service

RUN go test -v -race -tags live ./...
RUN CGO_ENABLED=0 go install -a /go/src/github.com/michaljirman/newsapp/apigateway-service/cmd/apigatewaysvc

FROM alpine:3.8
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/bin/apigatewaysvc .
CMD ["./apigatewaysvc"]

