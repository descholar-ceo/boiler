# getting base image
FROM golang

WORKDIR /boiler
COPY . /boiler
RUN go get github.com/mitchellh/go-homedir
CMD ["bin/boiler"]
