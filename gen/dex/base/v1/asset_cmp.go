package pb

func (a *Asset) Eq(o *Asset) bool {
	return a.Ticker == o.Ticker
}

func (a *DexAssetResponse) Eq(other *DexAssetResponse) bool {
	if len(a.A) != len(other.A) {
		return false
	}

	for _, a := range a.A {
		var has bool = false
		for _, b := range other.A {
			if a.Ticker == b.Ticker {
				has = true
				break
			}
		}
		if !has {
			return has
		}
	}

	return true
}
