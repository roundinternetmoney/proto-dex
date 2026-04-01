.PHONY: examples diff_iterator nats_asset_server nats_asset_client asset_compare protoc-nats pb buf all version tag release

LATEST_TAG := $(shell git describe --tags --abbrev=0 --match 'v[0-9]*' 2>/dev/null || echo v0.0.0)
VERSION := $(patsubst v%,%,$(LATEST_TAG))
PATCH_VERSION := $(shell echo $(VERSION) | awk -F. '{printf "%d.%d.%d", $$1, $$2, $$3+1}')
NEW_VERSION ?= $(PATCH_VERSION)

examples: diff_iterator nats_asset_server nats_asset_client asset_compare

asset_compare:
	go build -o bin/asset_compare examples/asset_compare/main.go

diff_iterator:
	go build -o bin/diff_iterator examples/diff_iterator/main.go

nats_asset_server:
	go build -o bin/nats_asset_server examples/nats_asset_server/main.go

nats_asset_client:
	go build -o bin/nats_asset_client examples/nats_asset_client/main.go


buf:
	buf dep update
	buf lint
	buf build

generate:
	buf generate

pb:
	go mod tidy
	go fmt ./...
	go vet ./...
	go test ./...
	go build ./...

protoc-nats:
#	latest version build command, but chucked streaming is currently unused and breaks imports on generation. 
# 	cd protoc-nats/extensions/proto && \ 
# 	buf generate --template buf.gen.yaml --path natsmicro/options.proto
	cd protoc-nats/ && \
	buf generate --template buf.gen.yaml --path extensions/proto/natsmicro/options.proto
	cp protoc-nats/gen/nats/micro/options.pb.go protoc-nats/tools/protoc-gen-nats-micro/nats/micro/options.pb.go
	cd protoc-nats && \
	go build -o tools/protoc-gen-nats-micro/protoc-gen-nats-micro ./tools/protoc-gen-nats-micro/

all: buf protoc-nats generate pb examples

push: buf
	buf push

version:
	@echo "Current version: $(VERSION)"
	@echo "Release version: $(NEW_VERSION)"

tag:
	git tag -a v$(NEW_VERSION) -m "Release v$(NEW_VERSION)"
	git push origin v$(NEW_VERSION)

release: tag
	gh release create v$(NEW_VERSION) \
		--title "v$(NEW_VERSION)" \
		--notes "Release v$(NEW_VERSION)"