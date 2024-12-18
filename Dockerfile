FROM golang:1.23.4-alpine AS build-stage
WORKDIR /app
COPY go.* .
COPY *.go .
RUN go build -o main .

FROM alpine:3.18 AS run-stage
WORKDIR /app
RUN adduser -D appuser
USER appuser
COPY --from=build-stage /app/main .
EXPOSE 1234
CMD ["./main"]
