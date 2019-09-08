package tax

import (
	t "github.com/kerscher/toy-tax-calculator/pkg/tax"
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

var (
	DefaultCurrency = accounting.DefaultAccounting("£", 2)

	// MaxIncome is used as a (disregarded) upper bound on last tax bands. The underlying calculation ignores it, but Go has no native concept of nullable types and we give this value so the type checks.
	MaxIncome = t.Amount{
		Value:    decimal.RequireFromString("10000000000000000"),
		Currency: DefaultCurrency,
	}
)

var Rates = map[string]t.Year{
	"2019": t.Year{
		Year: 2019,
		Allowance: t.Allowance{
			Value:    decimal.NewFromFloat(0),
			Currency: DefaultCurrency,
		},
		Bands: []t.Band{
			t.Band{
				Description: "Starter rate",
				Rate:        t.Rate(decimal.NewFromFloat(0.19)),
				Capacity: t.Amount{
					Value:    decimal.NewFromFloat(2000),
					Currency: DefaultCurrency,
				},
			},
			t.Band{
				Description: "Basic rate",
				Rate:        t.Rate(decimal.NewFromFloat(0.20)),
				Capacity: t.Amount{
					Value:    decimal.NewFromFloat(10149),
					Currency: DefaultCurrency,
				},
			},
			t.Band{
				Description: "Intermediate rate",
				Rate:        t.Rate(decimal.NewFromFloat(0.21)),
				Capacity: t.Amount{
					Value:    decimal.NewFromFloat(19429),
					Currency: DefaultCurrency,
				},
			},
			t.Band{
				Description: "Higher rate",
				Rate:        t.Rate(decimal.NewFromFloat(0.40)),
				Capacity: t.Amount{
					Value:    decimal.NewFromFloat(118419),
					Currency: DefaultCurrency,
				},
			},
			t.Band{
				Description: "Top rate",
				Rate:        t.Rate(decimal.NewFromFloat(0.46)),
				Capacity:    MaxIncome,
			},
		},
	},
}