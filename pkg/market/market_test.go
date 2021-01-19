package market

import (
	"reflect"
	"testing"
)

func TestListMarket_Strings(t *testing.T) {
	markets := ListMarket{Nasdaq, Kospi}
	expect := []string{"nasdaq", "kospi"}

	if !reflect.DeepEqual(markets.Strings(), expect) {
		t.Fatal("market strings result is not same with expected")
	}
}