package market

type (
	groww struct {
		id         string
		isTracking bool
		isLoggedIn bool
		// Bunch of kite related fields
	}
)

// Impliment all kite-market related functions
func (g *groww) Name() string {
	return "groww"
}

func (g *groww) GetAllSymbols() any {
	return "groww nahi hua"
}

func FetchGrowwMarket() {
	g := &groww{
		id:         "Na mera, na tera groww",
		isTracking: true,
		isLoggedIn: true,
	}

	FetchCurrentMarketName(g)

}
