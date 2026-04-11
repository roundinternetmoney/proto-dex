# AGENTS.md

## Cursor Cloud specific instructions

This is a Go + Protobuf library (not a deployable application). It defines DEX data models and generates Go code from `.proto` files using Buf and a custom `protoc-gen-nats-micro` plugin (in the `protoc-nats/` git submodule).

### Prerequisites

- **Go 1.25+** (already available via `go` toolchain auto-download)
- **Buf CLI** (`buf`) — must be installed at `/usr/local/bin/buf`
- **protoc-gen-go** — installed via `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`; lives in `~/go/bin/`
- **PATH** must include `~/go/bin` (already configured in `~/.bashrc`)

### Key commands

All development commands are in the `Makefile`. The main entry point is:

```
make all    # buf lint/build → build protoc-nats plugin → generate Go code → go vet/test/build → build examples
```

Individual targets: `make buf`, `make protoc-nats`, `make generate`, `make pb`, `make examples`.

### Caveats

- The `protoc-nats/` directory is a **git submodule**. Run `git submodule update --init --recursive` after cloning.
- `buf generate` uses a **local plugin** at `./protoc-nats/tools/protoc-gen-nats-micro/protoc-gen-nats-micro`. The `make protoc-nats` target must run before `make generate` to build this plugin.
- `make pb` runs `go fmt ./...` which may reformat generated files (`reader_nats.pb.go`, `shared_nats.pb.go`). This is expected and should not produce a dirty working tree after a clean `make all`.
- There are **no test files** in this repo. `go test ./...` succeeds with `[no test files]` for all packages.
- The NATS examples (`nats_asset_server`, `nats_asset_client`) require a running NATS server at `nats://localhost:4222`. The standalone examples (`diff_iterator`, `asset_compare`) work without any external services.
