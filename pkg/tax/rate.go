package tax

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Rate is a percentage between 0% and 100% expressed as a decimal.Decimal
type Rate decimal.Decimal

func NewRate(v decimal.Decimal) (Rate, error) {
	if v.GreaterThan(decimal.NewFromFloat(1)) || v.LessThan(decimal.NewFromFloat(0)) {
		return Rate(decimal.NewFromFloat(0)), fmt.Errorf("Can't convert %s to a rate. Value must be between 0 and 1", v)
	}
	return Rate(v), nil
}
