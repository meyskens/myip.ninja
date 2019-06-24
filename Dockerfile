FROM golang:1.12

ADD ./ /go/src/github.com/meyskens/myip.ninja

RUN cd /go/src/github.com/meyskens/myip.ninja && go install

CMD /go/bin/myip.ninja