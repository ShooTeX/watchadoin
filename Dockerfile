FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN go build -o watchadoin .

ENTRYPOINT ["/app/watchadoin"]
