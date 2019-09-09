package tax

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestRatesAreValidOrError(t *testing.T) {
	ts := []struct {
		in   decimal.Decimal
		want Rate
	}{
		{
			in:   decimal.NewFromFloat(0),
			want: Rate(decimal.NewFromFloat(0)),
		},
		{
			in:   decimal.NewFromFloat(0.5),
			want: Rate(decimal.NewFromFloat(0.5)),
		},
		{
			in:   decimal.NewFromFloat(32),
			want: Rate{},
		},
		{
			in:   decimal.NewFromFloat(-2389),
			want: Rate{},
		},
	}
	for _, c := range ts {
		got, _ := NewRate(c.in)
		if !decimal.Decimal(got).Equal(decimal.Decimal(c.want)) {
			t.Errorf("NewRate(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
