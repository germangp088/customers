FROM golang

WORKDIR /go/src/github.com/germangp088/customers/
COPY . .

RUN go get -d -v  github.com/gorilla/mux
RUN go install -v  github.com/gorilla/mux

RUN go build -o customers . 

ENTRYPOINT /go/bin/customers

EXPOSE 8080