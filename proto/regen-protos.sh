#!/usr/bin/env bash
set -x
set -e

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)
pushd "${SCRIPT_DIR}"

GOPATH=${GOPATH:=$HOME/go}
GOBIN=${GOBIN:=$GOPATH/bin}

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/mitchellh/protoc-gen-go-json
go get -u github.com/spinnaker/kork

if [[ ":$PATH:" != *":${GOBIN}:"* ]]; then
	echo "Adding $GOBIN to PATH"
	export PATH=$PATH:$GOBIN
fi

KORK_VERSION=`grep -oP "(?<=kork ).*(?= //)" $SCRIPT_DIR/../go.mod`

./protoc --go_out=paths=source_relative:$SCRIPT_DIR/stats/ \
  --go-json_out=$SCRIPT_DIR/stats/ \
  --proto_path=$GOPATH/pkg/mod/github.com/spinnaker/kork@$KORK_VERSION/kork-proto/src/main/proto/stats/ \
  $GOPATH/pkg/mod/github.com/spinnaker/kork@$KORK_VERSION/kork-proto/src/main/proto/stats/*.proto
popd
