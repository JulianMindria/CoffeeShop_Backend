FROM golang:1.18.10-alpine

WORKDIR /goapp

COPY . .

RUN go mod download
RUN go build -v -o /goapp/backendgo ./cmd/main.go

EXPOSE 8081

ENTRYPOINT [ "/goapp/backendgo" ]