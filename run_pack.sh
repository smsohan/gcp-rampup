set -e
set -x

pack config default-builder cnbs/sample-builder:bionic
pack build test-go-app --path . --buildpack ./go-buildpack


# pack buildpack package smsohan/go-cnb --config ./package.toml
# pack buildpack package smsohan/go-cnb --config ./package.toml --publish
# pack buildpack package go-buildpack.cnb --config ./package.toml --format file

# pack buildpack register smsohan/go-cnb

docker run -p 8080:8080 test-go-app
