FROM golang:1.9

ADD ./ /go/src/github.com/meyskens/myip.ninja

RUN cd /go/src/github.com/meyskens/myip.ninja && go install

CMD /go/bin/myip.ninja