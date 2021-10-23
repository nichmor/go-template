
FROM golang:1.17.1-alpine 
EXPOSE 8000

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH  

RUN mkdir -p $GOPATH/src/go-template
ADD . $GOPATH/src/go-template

WORKDIR $GOPATH/src/go-template 
RUN go build -o go-template . 

CMD ["/go/src/go-template/go-template"]