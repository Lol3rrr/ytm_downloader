FROM golang:latest

WORKDIR /go/src/ytm_downloader
COPY . .

RUN go get -d ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["ytm_downloader"]
