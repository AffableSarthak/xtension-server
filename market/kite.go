package market

import "github.com/google/uuid"

type (
	Kite struct {
		Id         string
		IsTracking bool
		IsLoggedIn bool
		// Bunch of Kite related fields
	}
)

// Implement all kite-market related functions
func (k *Kite) Name() string {
	_ = k.Id
	_ = k.IsLoggedIn
	_ = k.IsTracking
	return "kite"
}

func (k *Kite) GetAllSymbols() any {
	return make([]string, 5)
}

func NewKiteMarket() Kite {
	k := Kite{
		Id:         uuid.NewString(),
		IsTracking: true,
		IsLoggedIn: false,
	}

	// FetchCurrentMarketName(k)
	return k
}
