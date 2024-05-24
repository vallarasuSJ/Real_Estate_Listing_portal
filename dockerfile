FROM golang:1.22-bullseye
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app main.go
# ENTRYPOINT ["/app/app"]



FROM debian:bullseye-slim
WORKDIR /root
COPY --from=0 /app/app ./
COPY --from=0 /app/docker.json ./

EXPOSE 3000
ENTRYPOINT ["./app"]