set -e
set -x

pack config default-builder cnbs/sample-builder:bionic
pack build test-go-app --path . --buildpack ./go-buildpack
docker run -p 8080:8080 test-go-app
