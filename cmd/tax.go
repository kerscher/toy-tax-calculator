package cmd

import (
	"fmt"

	taxInternal "github.com/kerscher/toy-tax-calculator/internal/app"
	"github.com/kerscher/toy-tax-calculator/pkg/tax"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

func calculateTax(cmd *cobra.Command, args []string) {
	cfg, err := processArgs(cmd, args)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if y, ok := taxInternal.Rates[cfg.taxYear]; ok {
		giv, err := decimal.NewFromString(cfg.grossIncome)
		if err != nil {
			fmt.Printf("Gross income provided is invalid. Please verify your formatting. You used: %v", cfg.grossIncome)
			return
		}
		gi := tax.GrossIncome{
			Value:    giv,
			Currency: taxInternal.DefaultCurrency,
		}
		py, err := tax.NewPayerYear(y, gi)
		if err != nil {
			fmt.Printf("Could not calculate taxes for your income this given year. This is likely a bug. Report to the developers sending the error below:\n\n%v", err)
		}
		fmt.Print(py.String())
		return
	}
	fmt.Printf("Invalid year. Valid choices would have been:\n%v\n", intKeys(taxInternal.Rates))
}

func intKeys(m map[int]tax.Year) []int {
	ks := make([]int, len(m))
	i := 0
	for k := range m {
		ks[i] = k
		i++
	}
	return ks
}
