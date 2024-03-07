package market

import (
	"fmt"
	"reflect"
)

type (
	Market interface {
		Name() string
		GetAllSymbols() any
	}

	MarketProvider struct {
		m Market
	}

	MarketType interface {
		Kite | Growww
	}
)

func FetchCurrentMarket(m Market) {
	fmt.Println(m)
	fmt.Println(m.Name())
	fmt.Println(m.GetAllSymbols())
}

func FetchCurrentMarketName(m Market) {
	fmt.Println(reflect.TypeOf(m).Name(), "Market:", m.Name())
}

// Pattern for Interfaces
func NewMarketProvider[T Kite | Growww](m T) *T {
	return &m
}

func (mp *MarketProvider) Name() string {
	return "check kare mp pattern"
}

func (mp *MarketProvider) GetAllSymbols() any {
	return "checked, it works"
}
