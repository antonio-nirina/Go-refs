FROM golang:1.16-alpine
RUN apk add git
WORKDIR /server

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o /main

EXPOSE 8080
RUN chmod +x /main
CMD [ "/main" ]
