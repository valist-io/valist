FROM golang:1.16
WORKDIR /go/src/github.com/valist-io/valist/
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash - \
	&& apt install -yy nodejs
COPY . ./
ENV CGO_ENABLED 0
RUN make

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/valist-io/valist/valist ./
ENTRYPOINT ["./valist"]
CMD ["daemon"]