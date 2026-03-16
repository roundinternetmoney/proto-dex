Public protobuf spec for decentralised exchanges.

## CI/CD

This repo is set up for Buff Schema Registry:

- `Buf CI` action validates and registers proto changes on pull requests and pushes.


Consumers can depend directly on the [Buf-generated modues](https://buf.build/round-internet-money/dex/sdks/main:protobuf), for example in golang:

```bash
go get buf.build/gen/go/round-internet-money/dex/protocolbuffers/go@v1.36.11-20260316143239-8152bb253370.1
```

That is the primary distribution path. The separate generated-code repository is optional and acts as a library for consumers that do not want to depend on their own implementations.
