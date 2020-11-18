#!/usr/bin/env bash

if [[ ! -f "resources/openapi.yaml" ]]; then
	wget https://instana.github.io/openapi/openapi.yaml -O resources/openapi.yaml
fi

if [[ ! -f openapi-generator-cli.jar ]]; then
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.1/openapi-generator-cli-4.3.1.jar -O openapi-generator-cli.jar
fi

git rm -r -q api docs *.go

GO_POST_PROCESS_FILE="gofmt -s -w" java -jar openapi-generator-cli.jar generate -i resources/openapi.yaml -g go \
    -o instana \
    --git-repo-id go-instana \
    --git-user-id cedricziel \
    --additional-properties packageName=instana \
    --additional-properties isGoSubmodule=true \
    --additional-properties packageVersion=1.190.696 \
    --type-mappings=object=interface{} \
    --enable-post-process-file \
    --skip-validate-spec

git add .
