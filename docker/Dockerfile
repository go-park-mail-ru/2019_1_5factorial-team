FROM golang:alpine

WORKDIR /go/src/github.com/go-park-mail-ru/2019_1_5factorial-team

COPY . .

RUN go install github.com/go-park-mail-ru/2019_1_5factorial-team

ENTRYPOINT /go/bin/2019_1_5factorial-team

CMD ["main"]