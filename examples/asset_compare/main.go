package main

import (
	"fmt"

	basev1 "roundinternet.money/proto-dex/gen/dex/base/v1"
)

func main() {
	one := &basev1.Asset{
		Id:          "BTC-USD",
		Ticker:      "BTC",
		MinQuantity: 0.001,
		MaxQuantity: 100,
		LotSize:     0.001,
		TickSize:    0.1,
	}

	two := &basev1.Asset{
		Id:          "BTC-USD",
		Ticker:      "BTC",
		MinQuantity: 0.001,
		MaxQuantity: 100,
		LotSize:     0.001,
		TickSize:    0.1,
	}

	fmt.Printf("Assets equal: %v\n", one.Eq(two))

	three := &basev1.DexAssetResponse{
		A: []*basev1.Asset{one, one, one, one},
	}
	four := &basev1.DexAssetResponse{
		A: []*basev1.Asset{two, two, two, two},
	}

	fmt.Printf("Asset lists equal: %v\n", three.Eq(four))
}
