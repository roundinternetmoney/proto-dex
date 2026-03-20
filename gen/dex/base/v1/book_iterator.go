package pb

func NewBookDiff() *BookDiffStream {
	return &BookDiffStream{
		Asks: make([]*BookLevel, 0),
		Bids: make([]*BookLevel, 0),
	}
}

func NewBookLevel(p, s string) *BookLevel {
	return &BookLevel{Price: p, Size: s}
}

func (d *BookDiffStream) BidLevels() func(func(string, string) bool) {
	return func(yield func(string, string) bool) {
		for i := len(d.Bids) - 1; i >= 0; i-- {
			if !yield(d.Bids[i].Price, d.Bids[i].Size) {
				return
			}
		}
	}
}

func (d *BookDiffStream) AskLevels() func(func(string, string) bool) {
	return func(yield func(string, string) bool) {
		for i := len(d.Asks) - 1; i >= 0; i-- {
			if !yield(d.Asks[i].Price, d.Asks[i].Size) {
				return
			}
		}
	}
}
