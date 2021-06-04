FROM golang

COPY . /src
WORKDIR /src

CMD go test .
