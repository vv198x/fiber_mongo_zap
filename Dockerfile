FROM golang:1.19.3-alpine as build

WORKDIR /usr/local/go/src/fiber_mongo_zap/
ADD . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -v -a -o fiber_mongo_zap -ldflags="-w -s"

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/go/src/fiber_mongo_zap/fiber_mongo_zap /

ENTRYPOINT ["./fiber_mongo_zap"]