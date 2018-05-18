FROM golang:1-alpine

ADD . /tmp/generator-go

RUN apk --no-cache add git make nodejs-npm

RUN addgroup yo && \
    adduser -D -G yo yo

RUN npm install --global yo && \
    npm install --global /tmp/generator-go && \
    go get -u github.com/golang/dep/cmd/dep

USER yo

RUN mkdir /home/yo/go
ENV GOPATH "/home/yo/go"

CMD [ "yo", "go" ]
