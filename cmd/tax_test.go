package cmd

import (
	"testing"

	taxInternal "github.com/kerscher/toy-tax-calculator/internal/app"
	"github.com/kerscher/toy-tax-calculator/pkg/tax"
	"github.com/shopspring/decimal"
)

func TestTaxMatchesExample(t *testing.T) {
	inGrossIncome := tax.GrossIncome{
		Value:    decimal.NewFromFloat(43_500),
		Currency: taxInternal.DefaultCurrency,
	}
	inYear := 2019
	if inTaxYear, ok := taxInternal.Rates[inYear]; ok {
		want := tax.PayerYear{
			Year:        2019,
			Allowance:   inTaxYear.Allowance,
			GrossIncome: inGrossIncome,
			TaxableIncome: tax.TaxableIncome{
				Value:    decimal.NewFromFloat(31_650),
				Currency: taxInternal.DefaultCurrency,
			},
			Due: tax.Amount{
				Value:    decimal.NewFromFloat(6_518.69),
				Currency: taxInternal.DefaultCurrency,
			},
			PayerBands: []tax.PayerBand{
				tax.PayerBand{
					Band: tax.Band{
						Description: "Starter rate",
						Rate:        tax.Rate(decimal.NewFromFloat(0.19)),
						Capacity: tax.Amount{
							Value:    decimal.NewFromFloat(2_000),
							Currency: taxInternal.DefaultCurrency,
						},
					},
					Due: tax.Amount{
						Value:    decimal.NewFromFloat(380),
						Currency: taxInternal.DefaultCurrency,
					},
				},
				tax.PayerBand{
					Band: tax.Band{
						Description: "Basic rate",
						Rate:        tax.Rate(decimal.NewFromFloat(0.20)),
						Capacity: tax.Amount{
							Value:    decimal.NewFromFloat(10_149),
							Currency: taxInternal.DefaultCurrency,
						},
					},
					Due: tax.Amount{
						Value:    decimal.NewFromFloat(2_029.80),
						Currency: taxInternal.DefaultCurrency,
					},
				},
				tax.PayerBand{
					Band: tax.Band{
						Description: "Intermediate rate",
						Rate:        tax.Rate(decimal.NewFromFloat(0.21)),
						Capacity: tax.Amount{
							Value:    decimal.NewFromFloat(19_429),
							Currency: taxInternal.DefaultCurrency,
						},
					},
					Due: tax.Amount{
						Value:    decimal.NewFromFloat(4_080.09),
						Currency: taxInternal.DefaultCurrency,
					},
				},
				tax.PayerBand{
					Band: tax.Band{
						Description: "Higher rate",
						Rate:        tax.Rate(decimal.NewFromFloat(0.40)),
						Capacity: tax.Amount{
							Value:    decimal.NewFromFloat(118_419),
							Currency: taxInternal.DefaultCurrency,
						},
					},
					Due: tax.Amount{
						Value:    decimal.NewFromFloat(28.80),
						Currency: taxInternal.DefaultCurrency,
					},
				},
				tax.PayerBand{
					Band: tax.Band{
						Description: "Top rate",
						Rate:        tax.Rate(decimal.NewFromFloat(0.46)),
						Capacity:    taxInternal.MaxIncome,
					},
					Due: tax.Amount{
						Value:    decimal.NewFromFloat(0),
						Currency: taxInternal.DefaultCurrency,
					},
				},
			},
		}
		py, _ := tax.NewPayerYear(inTaxYear, inGrossIncome)
		// The condition below should use reflect.DeepEqual, but
		// despite having identical values DeepEqual is unable to
		// maintain equality. This is likely due to limitations with
		// how it handles pointer dereferencing for values inside
		// decimal.Decimal. To circumvent this we use the String
		// method for both.
		if py.String() != want.String() {
			t.Errorf("Calculated tax does not match reference.\n\nReference:\n----------\n%v\n\nCalculated:\n-----------\n %v", py, want)
			return
		}
		return
	}
	t.Errorf("Year %v rates does not exist for integration tests. Please rectify the internal library code.", inYear)
}
