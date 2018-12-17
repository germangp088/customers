FROM golang

# GOLANG
WORKDIR /go/src/github.com/germangp088/customers/
COPY . .

RUN go get -d -v  github.com/gorilla/mux
RUN go install -v  github.com/gorilla/mux

RUN go build .
RUN go install

# SWAGGER
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger
RUN swagger generate spec -o ./swagger.json --scan-models
RUN swagger serve -F=swagger swagger.json

ENTRYPOINT /go/bin/customers

EXPOSE 8080