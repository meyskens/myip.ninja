FROM golang:1.16

ADD ./ /go/src/github.com/meyskens/myip.ninja

RUN cd /go/src/github.com/meyskens/myip.ninja && go install

WORKDIR /

CMD /go/bin/myip.ninja