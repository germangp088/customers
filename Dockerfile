FROM golang

WORKDIR /go/src/app
COPY . .

ADD . /go/src/app/routes

RUN go get -d -v  github.com/gorilla/mux
RUN go install -v  github.com/gorilla/mux

RUN go build -o main . 

ENTRYPOINT /go/bin/main

EXPOSE 8080