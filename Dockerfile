FROM golang:1-alpine AS builder

WORKDIR /go/src/github.com/jyap808/cryptoEtfScrape

COPY . .

RUN go build -ldflags="-s -w"

FROM alpine:3

# Install the tzdata package to include timezone data files
RUN addgroup -S julian -g 1000 && \
    adduser -S julian -G julian -u 1000 && \
    apk add --no-cache tzdata

# Set the timezone
ENV TZ=America/Los_Angeles

COPY --from=builder /go/src/github.com/jyap808/cryptoEtfScrape/cryptoEtfScrape /usr/local/bin

USER julian

CMD ["/usr/local/bin/cryptoEtfScrape"]
