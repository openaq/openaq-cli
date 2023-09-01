package cmd

import (
	"github.com/spf13/cobra"
)

var Pretty bool
var JSON bool
var CSV bool
var Mini bool
var Limit int64
var Page int64
var FromDate string
var ToDate string
var PeriodName string
var ParameterType string
var IsoCode string

var countriesIDs []int64
var providersIDs []int64
var parametersIDs []int64

func addCountries(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64SliceVar(&countriesIDs, "countries", []int64{}, "filter results by country ID(s)")
}

func addParameters(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64SliceVar(&parametersIDs, "parameters", []int64{}, "filter results by parameters ID(s)")
}

func addParametersType(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&ParameterType, "type", "", "filter parameters to either `pollutants` or `meteorolgoical` parameters")
}

func addProviders(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64SliceVar(&providersIDs, "providers", []int64{}, "filter results by provider by provider ID(s)")
}

func addFormat(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVar(&JSON, "json", false, "output results as JSON")
	cmd.PersistentFlags().BoolVar(&CSV, "csv", false, "output results as CSV (Comma Separated Values)")
	cmd.PersistentFlags().BoolVar(&Pretty, "pretty", false, "pretty print")
	cmd.MarkFlagsMutuallyExclusive("json", "csv")
	cmd.MarkFlagsMutuallyExclusive("pretty", "csv")
}

func addLimit(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64Var(&Limit, "limit", 100, "limit")
}

func addPage(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64Var(&Page, "page", 1, "page")
}

func addFromDate(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&FromDate, "from", "2000-01-01", "from")
}

func addToDate(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&ToDate, "to", "", "to")
}

func addPeriodName(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&PeriodName, "period-name", "", "period-name")
}

func addMini(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVar(&Mini, "mini", false, "mini")
}

func addIsoCode(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&IsoCode, "iso", "", "Limit the results to a specific country using ISO 3166-1 alpha-2 code")
}
