FROM golang:alpine AS binary
ADD . /app
WORKDIR /app
RUN go build -o http

FROM alpine:latest
WORKDIR /app
COPY --from=binary /app/http /app
CMD ["/app/http"]
