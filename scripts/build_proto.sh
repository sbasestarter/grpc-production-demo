#!/bin/bash

if [[ ! -d proto ]]; then
    echo "terminate by wrong work directory"
    exit 1
fi

tempdir=$(mktemp -d grpc-production-demo-proto.XXXXXX)
protoc -I proto/ proto/hello.proto --go_out=plugins=grpc:${tempdir}
mv -f ${tempdir}/github.com/sbasestarter/grpc-production-demo.git/proto/gen/* proto/gen/go/
rm -rf ${tempdir}

