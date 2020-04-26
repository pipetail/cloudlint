FROM golang:1.14-alpine as builder

RUN set -eux \
	&& apk add --update --no-cache \
		make \
	&& mkdir -p /build

WORKDIR /build
ADD . /build/

RUN set -eux \
	&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

FROM scratch
COPY --from=builder /build/bin/cloudlint /usr/bin/cloudlint
ENTRYPOINT ["/usr/bin/cloudlint"]
