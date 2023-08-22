FROM golang:1.21.0-alpine AS binary
ADD . /app
WORKDIR /app
RUN go build -o whoami

FROM alpine:latest
WORKDIR /app
COPY --from=binary /app/whoami /app
CMD ["/app/whoami"]
