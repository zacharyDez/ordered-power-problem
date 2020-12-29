FROM golang:1.15

WORKDIR /go/src/
RUN go get github.com/stretchr/testify

COPY src problem 

CMD [""]