package market

type (
	Growww struct {
		Id        string
		IsRunning bool
		IsSane    bool
	}
)

// Impliment all kite-market related functions
func (g *Growww) Name() string {
	return "groww"
}

func (g *Growww) GetAllSymbols() any {
	return "groww nahi hua"
}

func FetchGrowwMarket() {
	g := &Growww{
		Id:        "Na mera, na tera groww",
		IsRunning: true,
		IsSane:    true,
	}

	FetchCurrentMarketName(g)

}
