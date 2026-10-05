FROM golang:1.27-bookworm

WORKDIR /app

COPY . .

EXPOSE 8080

CMD ["sh", "-c", "tail -f /dev/null"]

