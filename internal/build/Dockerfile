FROM golang:buster
WORKDIR /opt/build
COPY ./ ./
RUN go build -o ./dist/main testdata/main.go