package market

type (
	kite struct {
		id         string
		isTracking bool
		isLoggedIn bool
		// Bunch of kite related fields
	}
)

// Implement all kite-market related functions
func (k *kite) Name() string {
	_ = k.id
	_ = k.isLoggedIn
	_ = k.isTracking
	return "kite"
}

func (k *kite) GetAllSymbols() any {
	return make([]string, 5)
}

func FetchKiteMarket() {
	k := &kite{}

	FetchCurrentMarketName(k)

}
