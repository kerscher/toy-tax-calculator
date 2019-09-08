package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Cfg is a struct with configuration flags and arguments
type Cfg struct {
	taxYear int
	grossIncome string
}

var (
	cfg Cfg
)

var rootCmd = &cobra.Command{
	Use:   "toy-tax-calculator",
	Short: "Calculates (fictitious) taxes",
	Long:  `Calculates (fictitious) taxes based on gross income for a given year`,
	Run:   calculateTax,
}

// Execute runs once and is called by cobra for CLI processing
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVar(&cfg.taxYear, "tax-year", 2019, "Year where taxes apply")
	rootCmd.PersistentFlags().StringVar(&cfg.grossIncome, "gross-income", "0", "Total amount of gross income")
}

func processArgs(cmd *cobra.Command, args []string) (*Cfg, error) {
	y, err := cmd.Flags().GetInt("tax-year")
	if err != nil { return &Cfg{}, err }

	g, err := cmd.Flags().GetString("gross-income")
	if err != nil { return &Cfg{}, err }

	return &Cfg{taxYear: y, grossIncome: g}, nil
}
