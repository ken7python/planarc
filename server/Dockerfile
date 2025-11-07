FROM golang:1.25
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org,direct
RUN go mod download

COPY . .

CMD ["air", "-c", "./air.toml"]
# CMD ["sh", "./server.sh" ]
