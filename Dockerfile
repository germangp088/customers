# FROM golang:1.11-windowsservercore
FROM golang
# ENV GOPATH "C:\\Users\\gdario\\go"
# ENV GOROOT "C:\\Go"

# WORKDIR /Users/gdario/go/docker
# COPY . .
# RUN git clone https://github.com/gorilla/mux /go/src/github.com/gorilla/mux

WORKDIR /go/src/app
COPY . .

RUN go get -d -v  github.com/gorilla/mux
RUN go install -v  github.com/gorilla/mux

# RUN git clone github.com/gorilla/mux
# RUN go install github.com/gorilla/mux
# RUN go get -u github.com/gorilla/mux

# RUN mkdir /main 
# ADD . /main/ 
# WORKDIR /main 

RUN go build -o main . 

# CMD ["/main"]

EXPOSE 8080