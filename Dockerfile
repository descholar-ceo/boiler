# getting base image
FROM linuxmintd/mint19.3-amd64

WORKDIR /boiler
COPY . /boiler
CMD ["bin/main"]
