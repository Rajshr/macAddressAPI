# ------ Go Builder Image ------
FROM golang:1.13 AS builder

#------ Environment Variables ------
ENV MAC_ADDRESS=""
ENV MAC_API_KEY=""

WORKDIR /go/src/macapi
COPY . .
RUN go get
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM alpine:3.4

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/macapi /bin/

ENTRYPOINT ["/bin/macapi"]