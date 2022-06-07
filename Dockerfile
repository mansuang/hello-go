FROM golang:1.11

RUN mkdir /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app

#install the external dependecny that we have in the code
RUN go get -u github.com/gorilla/mux
#make the build of the app, output will be an executable file named main
RUN go build -o main .

CMD ["/go/src/app/main"]