# getting base image
FROM linuxmintd/mint19.3-amd64

WORKDIR /stock-picker
COPY . /stock-picker
CMD ["bin/main"]
