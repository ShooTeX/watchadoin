FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN \
  go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps && \
  go build -o watchadoin .

CMD ["/app/watchadoin"]
