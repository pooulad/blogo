FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main ./cmd/blogo

EXPOSE 8080

CMD ["/app/main"]
# CMD ["/app/main","--cfg","./config/config.json"]
# FROM golang:alpine as builder

# LABEL maintainer="Pooulad <poooooladi@gmail.com>"

# RUN apk update && apk add --no-cache git

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# WORKDIR /app/cmd/blogo

# RUN go build -o /app/blogo .

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# COPY --from=builder /app/blogo .
# COPY --from=builder /app/.env .       
# COPY --from=builder /app/config/config.json .

# COPY check-db.sh /root/check-db.sh
# RUN chmod +x /root/check-db.sh

# EXPOSE 8080

# CMD ["/root/check-db.sh"]
