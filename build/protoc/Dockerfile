FROM alpine:3.12.0 as curl
RUN apk add --no-cache curl unzip
RUN adduser --disabled-password --gecos '' curl
USER curl
WORKDIR /home/curl

FROM curl as protoc
ARG PROTOC_VERSION=3.9.0
RUN curl -vLo protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
  && unzip protoc.zip \
  # We want all files/folders in bin/ to be executable
  && chmod -R 755 bin/ \
  # Files under include/ should not be executable so we need
  # to run a separate command for files/folders
  && find include/ -type d -exec chmod 755 {} \+ \
  && find include/ -type f -exec chmod 644 {} \+

FROM curl as protoc-gen-go
ARG PROTOC_GEN_GO_TAG=v1.23.0
RUN curl -L https://github.com/protocolbuffers/protobuf-go/releases/download/$PROTOC_GEN_GO_TAG/protoc-gen-go.$PROTOC_GEN_GO_TAG.linux.amd64.tar.gz \
  | tar -xz \
  && chmod 755 protoc-gen-go

FROM golang:1.14.4-buster as protoc-gen-go-json
ARG PROTOC_GEN_GO_JSON_TAG=069933b8c8344593ed8905d46d59c6647c886f47
RUN adduser --disabled-password --gecos '' golang
USER golang
RUN go get -d github.com/mitchellh/protoc-gen-go-json \
 && git -C "$(go env GOPATH)"/src/github.com/mitchellh/protoc-gen-go-json checkout $PROTOC_GEN_GO_JSON_TAG \
 && go install github.com/mitchellh/protoc-gen-go-json \
 && chmod 755 /go/bin/protoc-gen-go-json

FROM curl as kork
ARG KORK_VERSION=7.45.9
RUN curl -L https://github.com/spinnaker/kork/archive/v$KORK_VERSION.tar.gz | tar -xz \
  && mv kork-$KORK_VERSION kork \
  && find kork/ -type d -exec chmod 755 {} \+ \
  && find kork/ -type f -exec chmod 644 {} \+

FROM debian:buster-slim
COPY --chown=root:root --from=protoc /home/curl/bin/protoc /usr/local/bin/protoc
COPY --chown=root:root --from=protoc /home/curl/include/google /usr/local/include/google
COPY --chown=root:root --from=protoc-gen-go /home/curl/protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --chown=root:root --from=protoc-gen-go-json /go/bin/protoc-gen-go-json /usr/local/bin/protoc-gen-go-json
COPY --chown=root:root --from=kork /home/curl/kork/kork-proto/src/main/proto/stats/ /opt/proto/
COPY genproto.sh /usr/local/bin/genproto.sh
RUN chmod 755 /usr/local/bin/genproto.sh
RUN adduser --disabled-password --gecos '' --no-create-home protoc
USER protoc
ENTRYPOINT ["/usr/local/bin/genproto.sh"]
