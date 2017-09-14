FROM alpine
MAINTAINER Samuel Cozannet <samnco@gmail.com>

COPY gopath/bin/demo-app /go/bin/demo-app

ENTRYPOINT /go/bin/demo-app
