FROM golang:1.16-alpine
WORKDIR /app
COPY . .
RUN go mod init local && go build -o app .
CMD ["./app", "2381"]
