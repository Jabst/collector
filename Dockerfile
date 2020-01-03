FROM golang

ADD . /go/src/consumer/collector/

RUN go build /go/src/consumer/collector/

COPY /config /go/config
RUN mv /go/collector /go/bin/

ENTRYPOINT /go/bin/collector

EXPOSE 80