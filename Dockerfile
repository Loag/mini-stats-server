FROM golang:alpine as builder

RUN apk update && apk add --no-cache ca-certificates upx && update-ca-certificates

WORKDIR /usr/local/src

COPY . .

RUN go mod download

# use linker flags to remove debug symbols and optimize build
# make sure cgo is off to build a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o dist cmd/app/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/src/dist .
# COPY --from=builder /usr/local/src/sql/ sql/

EXPOSE 5000

CMD ["/dist"]