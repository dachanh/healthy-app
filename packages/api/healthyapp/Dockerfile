FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
RUN chmod +x /app/start.sh
EXPOSE 8176
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
