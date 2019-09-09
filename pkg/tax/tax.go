// Package tax provides currency-independent progressive tax functionality
//
// Internally everything is kept as decimal.Decimal, but users are required to also choose a currency when initialising and using values.
package tax

import (
	"fmt"
	"strings"

	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

// Amount is a value in a given currency
type Amount struct {
	Value    decimal.Decimal
	Currency *accounting.Accounting
}

func (a Amount) String() string {
	return a.Currency.FormatMoneyDecimal(a.Value)
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

// Converts a Band into a PayerBand. This is unexported as the Amount passed to it must be enforced elsewhere, as nothing guarantees a user will keep invariants related to GrossIncome. The Amount returned is the residual Amount if it exceeds Capacity for that band.
func (b *Band) due(a Amount) (PayerBand, Amount, error) {
	if b.Capacity.Currency != a.Currency {
		return PayerBand{}, Amount{}, fmt.Errorf("Cannot calculate taxes due. Currencies for Band and income Amount are not identical.")
	}
	if a.Value.LessThanOrEqual(b.Capacity.Value) {
		return PayerBand{
				Band: *b,
				Due: Amount{
					Value:    a.Value.Mul(decimal.Decimal(b.Rate)),
					Currency: a.Currency,
				},
			}, Amount{
				Value:    decimal.NewFromFloat(0),
				Currency: a.Currency,
			}, nil
	}
	return PayerBand{
			Band: *b,
			Due: Amount{
				Value:    b.Capacity.Value.Mul(decimal.Decimal(b.Rate)),
				Currency: a.Currency,
			},
		}, Amount{
			Value:    a.Value.Sub(b.Capacity.Value),
			Currency: a.Currency,
		}, nil
}

func (pb PayerBand) String() string {
	return fmt.Sprintf("%v: %v", pb.Band.Description, Amount(pb.Due).String())
}

// A Year has an allowance value and an ordered sequence of TaxBand
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
func (g *GrossIncome) Taxable(a Allowance) (TaxableIncome, error) {
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

//

// PayerYear is a Year with a sequence of PayerBand instead of Band and payer income information
type PayerYear struct {
	Year          int
	Allowance     Allowance
	GrossIncome   GrossIncome
	TaxableIncome TaxableIncome
	Due           Amount
	PayerBands    []PayerBand
}

// Calculates taxes of a taxpayer given their gross income and a tax Year
func NewPayerYear(y Year, g GrossIncome) (PayerYear, error) {
	ti, err := g.Taxable(y.Allowance)
	if err != nil {
		return PayerYear{}, err
	}
	tti := ti
	d := Amount{Value: decimal.NewFromFloat(0), Currency: g.Currency}
	var pbs []PayerBand
	for _, b := range y.Bands {
		pb, _, err := b.due(Amount(tti))
		if err != nil {
			return PayerYear{}, err
		}
		ntv := tti.Value.Sub(b.Capacity.Value)
		if ntv.IsNegative() {
			ntv = decimal.NewFromFloat(0)
		}
		tti = TaxableIncome{
			Value:    ntv,
			Currency: tti.Currency,
		}
		d.Value = d.Value.Add(pb.Due.Value)
		pbs = append(pbs, pb)
	}
	return PayerYear{
		Year:          y.Year,
		Allowance:     y.Allowance,
		GrossIncome:   g,
		TaxableIncome: ti,
		Due:           d,
		PayerBands:    pbs,
	}, nil
}

func (p PayerYear) String() string {
	var pbstr strings.Builder
	for _, pb := range p.PayerBands {
		pbstr.WriteString(pb.String())
		pbstr.WriteString("\n")
	}
	return fmt.Sprintf(`
Tax year: %d

Gross salary: %s

Personal allowance: %s

Taxable income: %s

%s

Total tax due: %v
`,
		p.Year,
		Amount(p.GrossIncome).String(),
		Amount(p.Allowance).String(),
		Amount(p.TaxableIncome).String(),
		pbstr.String(),
		Amount(p.Due).String(),
	)
}
