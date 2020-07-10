FROM alpine:3.12.0 as curl
RUN apk add --no-cache curl unzip

FROM curl as protoc
ARG PROTOC_VERSION=3.9.0
RUN curl -vLo protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
 && unzip protoc.zip

FROM curl as protoc-gen-go
ARG PROTOC_GEN_GO_TAG=v1.23.0
RUN curl -L https://github.com/protocolbuffers/protobuf-go/releases/download/$PROTOC_GEN_GO_TAG/protoc-gen-go.$PROTOC_GEN_GO_TAG.linux.amd64.tar.gz \
  | tar -xz

FROM golang:1.14.4-buster as protoc-gen-go-json
ARG PROTOC_GEN_GO_JSON_TAG=069933b8c8344593ed8905d46d59c6647c886f47
RUN go get -d github.com/mitchellh/protoc-gen-go-json \
 && git -C "$(go env GOPATH)"/src/github.com/mitchellh/protoc-gen-go-json checkout $PROTOC_GEN_GO_JSON_TAG \
 && go install github.com/mitchellh/protoc-gen-go-json

FROM curl as kork
ARG KORK_VERSION=7.45.9
RUN curl -L https://github.com/spinnaker/kork/archive/v$KORK_VERSION.tar.gz | tar -xz \
  && mv kork-$KORK_VERSION kork

FROM debian:buster-slim
COPY --from=protoc /bin/protoc /usr/local/bin/protoc
COPY --from=protoc /include/google /usr/local/include/google
COPY --from=protoc-gen-go /protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --from=protoc-gen-go-json /go/bin/protoc-gen-go-json /bin/
COPY --from=kork /kork/kork-proto/src/main/proto/stats/ /opt/proto/
COPY genproto.sh .
ENTRYPOINT ["./genproto.sh"]
