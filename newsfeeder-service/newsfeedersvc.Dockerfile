FROM golang:1.13.4 as builder

COPY vendor /go/src/github.com/michaljirman/newsapp/vendor

RUN mkdir -p /go/src/github.com/michaljirman/newsapp/src/newsfeeder-service
COPY newsfeeder-service /go/src/github.com/michaljirman/newsapp/newsfeeder-service
WORKDIR /go/src/github.com/michaljirman/newsapp/newsfeeder-service

RUN go test -v -race -tags live ./...
RUN CGO_ENABLED=0 go install -a /go/src/github.com/michaljirman/newsapp/newsfeeder-service/cmd/newsfeedersvc

FROM alpine:3.8
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/bin/newsfeedersvc .
COPY --from=builder /go/src/github.com/michaljirman/newsapp/newsfeeder-service/migrations /migrations
CMD ["./newsfeedersvc"]

