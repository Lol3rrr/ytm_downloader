FROM lol3r/golang-ffmpeg:latest

WORKDIR /go/src/ytm_downloader
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80

CMD ["ytm_downloader"]
