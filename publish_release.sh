#!/bin/bash

if [[ -z "$1" ]]; then 
  echo "need tag/version in format v1.x.y"
  exit 1
else
  TAG=$1
fi

CGO_ENABLED=0 go build -o bin/sensu-http-ping-handler cmd/sensu-http-ping-handler/main.go
tar czf sensu-http-ping-handler_${TAG}_linux_amd64.tar.gz bin/

sha512sum sensu-http-ping-handler_${TAG}_linux_amd64.tar.gz > sensu-http-ping-handler_${TAG}_sha512_checksums.txt
SHA_HASH_ONLY=$(cut -d " " -f 1 sensu-http-ping-handler_${TAG}_sha512_checksums.txt)

sed "s/__TAG__/${TAG}/g" sensu/asset_template.tpl > sensu/asset.yaml
sed -i "s/__SHA__/${SHA_HASH_ONLY}/g" sensu/asset.yaml

mkdir -p artifacts
rm -f artifacts/*
mv sensu-http-ping-handler_${TAG}_linux_amd64.tar.gz sensu-http-ping-handler_${TAG}_sha512_checksums.txt artifacts/

git add .
git commit
git tag $TAG
git push && git push --tags
