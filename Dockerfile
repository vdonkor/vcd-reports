FROM golang:latest
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_HOME=/billing
WORKDIR $APP_HOME
RUN go get https://github.com/vdonkor/vcd-reports.git