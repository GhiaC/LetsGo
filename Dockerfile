FROM golang

WORKDIR $GOPATH/src/github.com/aoacademy/letsgo

ENV GO111MODULE=on

COPY . .

RUN go get -v ./...

RUN go install -v ./...

RUN go build main.go

CMD ["./main"]