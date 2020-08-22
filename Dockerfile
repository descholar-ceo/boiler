# getting base image
FROM linuxmintd/mint20-amd64

WORKDIR /boiler
COPY . /boiler
CMD ["bin/main"]
