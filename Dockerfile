
FROM golang:alpine

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

RUN go mod download

COPY . .

RUN go build -o ./main

EXPOSE 7080

CMD ["./main"]

	