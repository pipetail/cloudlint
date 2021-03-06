FROM golang:1.14-alpine as builder

RUN set -eux \
	&& apk add --update --no-cache \
		make \
		ca-certificates \
	&& mkdir -p /build

WORKDIR /build
ADD . /build/

RUN set -eux \
	&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

FROM alpine
COPY --from=builder /build/bin/cloudlint /usr/bin/cloudlint
ENV AWS_SDK_LOAD_CONFIG=1
ENTRYPOINT ["/usr/bin/cloudlint"]
