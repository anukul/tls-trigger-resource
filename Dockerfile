FROM      golang:alpine AS builder
WORKDIR   /tls-trigger-resource/
COPY      . .
RUN       go build -o /assets/check github.com/anukul/tls-trigger-resource/check/
RUN       cp ./in/in ./out/out /assets/

FROM      alpine
COPY      --from=builder /assets/ /opt/resource/
