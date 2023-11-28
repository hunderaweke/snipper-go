package internal

import "os"

func DockerWriter() {
	file, err := os.Create("Dockerfile")
	if err != nil {
		panic(err)
	}
	content := `
FROM golang:alpine

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

RUN go mod download

COPY . .

RUN go build -o ./main

EXPOSE 7080

CMD ["./main"]

	`
	file.WriteString(content)
	defer file.Close()
}

func DockerComposeWriter() {
	file, err := os.Create("docker-compose.yml")
	if err != nil {
		panic(err)
	}
	content := `services:
    mysql:
        image: mysql:8.2
        restart: always
        networks:
        - backend
        env_file:
        - ./.env
        volumes:
        - data:/var/lib/mysql
        ports:
        - 4060:3306
        environment:
        MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD}
        MYSQL_DATABASE: ${MYSQL_DATABASE}
        MYSQL_USER: ${MYSQL_USER}
        MYSQL_PASSWORD: ${MYSQL_PASSWORD}
        healthcheck:
            test: ["CMD", "mysqladmin" ,"ping", "-h", "localhost"]
            timeout: 20s
            retries: 10
    backend:
        build: ./backend
        ports:
        - 7080:7080
        env_file:
        - ./backend/.env
        networks:
        - backend
        depends_on:
            mysql:
        condition: service_healthy
networks:
    backend:
volumes:
    data:
`
	file.WriteString(content)
	defer file.Close()
}
