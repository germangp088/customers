FROM golang

WORKDIR /go/src/github.com/germangp088/customers/
COPY . .

RUN go get -d -v  github.com/gorilla/mux
RUN go install -v  github.com/gorilla/mux

RUN go build .
RUN go install

ENTRYPOINT /go/bin/customers

EXPOSE 8080