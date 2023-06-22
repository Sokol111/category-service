FROM golang:1.20.5-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o service ./cmd/main.go
#RUN chmod +x ./service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/service .
COPY --from=builder /app/config.yml .
EXPOSE 50052
CMD [ "/app/service" ]