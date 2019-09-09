package tax

import (
	"testing"

	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

var (
	testBand = Band{
		Description: "Testing rate",
		Rate:        Rate(decimal.NewFromFloat(0.5)),
		Capacity: Amount{
			Value:    decimal.NewFromFloat(1_000_000),
			Currency: accounting.DefaultAccounting("£", 2),
		},
	}
)

func TestDueIsZeroOnZeroAmount(t *testing.T) {
	in := Amount{
		Value:    decimal.NewFromFloat(0),
		Currency: testBand.Capacity.Currency,
	}
	pb, _, _ := testBand.due(in)
	if !decimal.Decimal(pb.Due.Value).Equal(decimal.NewFromFloat(0)) {
		t.Errorf("Zero income should lead to zero taxes due. Due: %v", pb.Due.Value)
	}
}

func TestDueIsCompatibleWithRate(t *testing.T) {
	in := Amount{
		Value:    decimal.NewFromFloat(500_000),
		Currency: testBand.Capacity.Currency,
	}
	want := decimal.NewFromFloat(250_000)
	pb, _, _ := testBand.due(in)
	if !decimal.Decimal(pb.Due.Value).Equal(want) {
		t.Errorf("Taxes due != (Income * Rate). Due: %v, Income: %v, Rate: %v, Expected: %v", pb.Due.Value, in, decimal.Decimal(pb.Rate).String(), want)
	}
}
