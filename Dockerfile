FROM golang
WORKDIR /app
EXPOSE 8080

COPY ["/", "/app"]
RUN go get ./
RUN GOOS=linux GOARCH=arm64 go build -o /tinywatcher
ENTRYPOINT [ "/tinywatcher"]