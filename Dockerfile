FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -i ./cmd/url-shorten/main.go -o ./build/url-shorten
FROM alpine
RUN adduser -S -D -H -h /api apiusr
USER appuser
COPY --from=builder /build/url-shorten /api/
WORKDIR /api
CMD ["./url-shorten"]

EXPOSE 3077
