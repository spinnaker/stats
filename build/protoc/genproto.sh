#!/bin/sh
set -e

PROTOC_FLAGS=""

gen_proto() {
  mkdir -p /tmp/staging/go
  find /opt/proto -name *.proto | sort | xargs protoc \
    $PROTOC_FLAGS \
    --proto_path=/opt/proto \
    --go_out=paths=source_relative:/tmp/staging/go \
    --go-json_out=tmp/staging/go
}

update_proto() {
  gen_proto
  rm -rf /mnt/output/go/*
  cp -r /tmp/staging/go/* /mnt/output/go/
}

check_proto() {
  gen_proto
  if ! diff -r /tmp/staging /mnt/output; then
    echo "Protocol buffer compilation out of date. Please run 'make proto' and commit the changes."
    echo "To help with debugging, the difference between the committed and generated code is printed above."
    exit 1
  fi
}

command=$1
case $command in
  update)
    update_proto
    ;;
  check)
    check_proto
    ;;
  *)
    echo "Unrecognized command: $command"
    exit 2
    ;;
esac
