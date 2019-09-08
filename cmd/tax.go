package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

func calculateTax(cmd *cobra.Command, args []string) {
	cfg, err := processArgs(cmd, args)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf(`This would have calculated your taxes for:
Tax year: %v

Gross Salary: %v

Personal allowance: <unimplemented>

Taxable income: <unimplemented>

Starter rate: <unimplemented>
Basic rate: <unimplemented>
Intermediate rate: <unimplemented>
Higher rate: <unimplemented>

Total tax due: <unimplemented>
`, cfg.taxYear, cfg.grossIncome)
}
