# Protobuf implementations for DEXs

[![Go Reference](https://pkg.go.dev/badge/roundinternet.money/proto.svg)](https://pkg.go.dev/roundinternet.money/proto)

## Getting started

- Requires Go 1.25+.
- Install from GitHub: `go get github.com/roundinternetmoney/pb-dex`

## Example Usage

```
package main

import (
	"fmt"

	"roundinternet.money/pb-dex"
)

func main() {
	diff := pb.NewBookDiff()
	diff.Bids = append(diff.Bids, pb.NewBookLevel("102", "2"))
	diff.Bids = append(diff.Bids, pb.NewBookLevel("101", "3"))

	for p, s := range diff.BidLevels() {
		fmt.Printf("BID | Price: %s Size: %s\n", p, s)
	}
}
```

For more complete usage examples see the [examples/](./examples/) folder in this repository.

## CI/CD

This repo is set up for Buff Schema Registry:

- `Buf CI` action validates and registers proto changes on pull requests and pushes.


`.proto` consumers can depend directly on the [Buf-generated modues](https://buf.build/round-internet-money/dex/sdks/main:protobuf)

Contributing
-------------
Contributions are welcome! Please open issues or pull requests as needed.