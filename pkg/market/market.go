package market

type Market string

type ListMarket []Market

func (lm *ListMarket) Strings() []string {
	var strings []string
	for _, m := range *lm {
		strings = append(strings, string(m))
	}
	return strings
}

var Nasdaq Market = "nasdaq"
var Kospi Market = "kospi"
var SupportedMarkets = ListMarket{Nasdaq, Kospi}
