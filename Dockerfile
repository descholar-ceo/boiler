# getting base image
FROM golang

WORKDIR /boiler
COPY . /boiler
CMD ["bin/main"]
