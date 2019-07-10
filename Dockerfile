FROM golang:1.11-alpine as builder

RUN grep nobody /etc/passwd > /etc/passwd.nobody \
    && grep nobody /etc/group > /etc/group.nobody
RUN apk --no-cache update \
    && apk add --no-cache ca-certificates git \
    && wget -O- https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . $GOPATH/src/github.com/GuilhermeFirmiano/mock-google-maps
WORKDIR $GOPATH/src/github.com/GuilhermeFirmiano/mock-google-maps

RUN dep ensure -v \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/group.nobody /etc/group
COPY --from=builder /etc/passwd.nobody /etc/passwd
USER nobody

COPY --from=builder /go/src/github.com/GuilhermeFirmiano/mock-google-maps/app .
COPY --from=builder /go/src/github.com/GuilhermeFirmiano/mock-google-maps/.env .

EXPOSE 9000
ENTRYPOINT ["/app"]