// Package tax provides currency-independent progressive tax functionality
//
// Internally everything is kept as decimal.Decimal, but users are required to also choose a currency when initialising and using values.
package tax

import (
	"fmt"

	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

// Amount is a value in a given currency
type Amount struct {
	Value    decimal.Decimal
	Currency *accounting.Accounting
}

// Allowance is the maximum tax-exempt income a taxpayer has
type Allowance Amount

// TaxBand consists of a description, a rate and a capacity. The actual band is only meaningful in the context of other TaxBands, as a sequence of them naturally leads to a minimum and a maximum for a particular band depending on its capacity.
//
// This particular design choice ensures bands are never overlapping and ensures correctness by construction.
type Band struct {
	Description string
	Rate        Rate
	Capacity    Amount
}

// A PayerBand is a Band enriched with the amount due by a given taxpayer
type PayerBand struct {
	Band
	Due Amount
}

// A Year has an allowance value and an ordered sequence of TaxBand consisting of the progressive taxation bands
type Year struct {
	Year      int
	Allowance Allowance
	Bands     []Band
}

// GrossIncome is the total income of a taxpayer before any taxes are calculated
type GrossIncome Amount

// TaxableIncome is the part of a taxpayer income due taxes in a given year
type TaxableIncome Amount

// Taxable converts GrossIncomes into TaxableIncomes
func (g GrossIncome) Taxable(a Allowance) (TaxableIncome, error) {
	if g.Currency != a.Currency {
		return TaxableIncome{}, fmt.Errorf("Cannot calculate taxable income. Currencies for GrossIncome and Allowance are not identical.")
	}
	t := g.Value.Sub(a.Value)
	if t.IsNegative() {
		return TaxableIncome{
			Value:    decimal.NewFromFloat(0),
			Currency: g.Currency,
		}, nil
	}
	return TaxableIncome{
		Value:    t,
		Currency: g.Currency,
	}, nil
}

// PayerYear is a Year with a sequence of PayerBand instead of Band and payer income information
type PayerYear struct {
	Year          int
	Allowance     Allowance
	GrossIncome   GrossIncome
	TaxableIncome TaxableIncome
	PayerBands    []PayerBand
}

func NewPayerYear(y Year, g GrossIncome) (PayerYear, error) {
	return PayerYear{}, fmt.Errorf("Calculating taxes is currently unimplemented!")
}

func (p *PayerYear) String() string {
	return "<unimplemented>"
}
