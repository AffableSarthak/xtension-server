package market

import (
	"fmt"
	"reflect"
)

type (
	market interface {
		Name() string
		GetAllSymbols() any
	}

	marketProvider struct {
		m market
	}
)

func FetchCurrentMarket(m market) {
	fmt.Println(m)
	fmt.Println(m.Name())
	fmt.Println(m.GetAllSymbols())
}

func FetchCurrentMarketName(m market) {
	fmt.Println(reflect.TypeOf(m).Name(), "Market:", m.Name())
}

// Pattern for Interfaces
func NewMarketProvider(m market) *marketProvider {
	return &marketProvider{
		m: m,
	}
}

func (mp marketProvider) Name() string {
	return "check kare mp pattern"
}

func (mp marketProvider) GetAllSymbols() any {
	return "checked, it works"
}

// market.FetchGrowwMarket()
// market.FetchKiteMarket()

// type T struct {
// 	A int
// 	B string
// }
// t := T{23, "skidoo"}
// s := reflect.ValueOf(&t).Elem()
// typeOfT := s.Type()
// s.NumField()
// for i := 0; i < s.NumField(); i++ {
// 	f := s.Field(i)
// 	fmt.Printf("%d: %s %s = %v\n", i,
// 		typeOfT.Field(i).Name, f.Type(), f.Interface())
// }
