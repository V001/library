FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
EXPOSE 8000
CMD [ "/app/library" ]
RUN go build -o library ./cmd
