FROM golang:1.14

WORKDIR /
COPY . .
RUN go mod download
EXPOSE 5000

CMD ["go","run","main.go"]