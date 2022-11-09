FROM golang:1.18.8-alpine3.16 as builder
RUN mkdir /build
ADD go.mod *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o blockchain-go .

FROM alpine:3.16.2
COPY --from=builder /build/blockchain-go .
ENTRYPOINT [ "./blockchain-go" ]