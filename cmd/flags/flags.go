package flags

import (
	"github.com/spf13/cobra"
)

var pretty bool
var json bool
var csv bool
var mini bool
var limit int64
var page int64
var fromDate string
var toDate string
var periodName string
var parameterType string
var isoCode string
var radius int64
var coordinates []float64
var bbox []float64

var countriesIDs []int64
var providersIDs []int64
var parametersIDs []int64

func AddCountries(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64SliceVar(&countriesIDs, "countries", []int64{}, "filter results by country ID(s)")
}

func AddParameters(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64SliceVar(&parametersIDs, "parameters", []int64{}, "filter results by parameters ID(s)")
}

func AddParametersType(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&parameterType, "type", "", "filter parameters to either `pollutants` or `meteorolgoical` parameters")
}

func AddProviders(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64SliceVar(&providersIDs, "providers", []int64{}, "filter results by provider by provider ID(s)")
}

func AddFormat(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVar(&json, "json", false, "output results as JSON")
	cmd.PersistentFlags().BoolVar(&csv, "csv", false, "output results as CSV (Comma Separated Values)")
	cmd.PersistentFlags().BoolVar(&pretty, "pretty", false, "pretty print")
	cmd.MarkFlagsMutuallyExclusive("json", "csv")
	cmd.MarkFlagsMutuallyExclusive("pretty", "csv")
}

func AddLimit(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64Var(&limit, "limit", 100, "limit")
}

func AddPage(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64Var(&page, "page", 1, "page")
}

func AddFromDate(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&fromDate, "from", "2000-01-01", "from")
}

func AddToDate(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&toDate, "to", "", "to")
}

func AddPeriodName(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&periodName, "period-name", "", "period-name")
}

func AddMini(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVar(&mini, "mini", false, "mini")
}

func AddIsoCode(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&isoCode, "iso", "", "Limit the results to a specific country using ISO 3166-1 alpha-2 code")
}

func AddRadiusSearch(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64Var(&radius, "radius", 0, "distance in meters to search around `coordinates`")
	cmd.PersistentFlags().Float64SliceVar(&coordinates, "coordinates", nil, "Coordinate pair of center point to perform radius search. In form latitude,longitude i.e. y,x")
	cmd.MarkFlagsRequiredTogether("radius", "coordinates")
}

func AddBBox(cmd *cobra.Command) {
	cmd.PersistentFlags().Float64SliceVar(&bbox, "bbox", nil, "A bounding box to search within in form minx,miny,maxx,maxy")
	cmd.MarkFlagsMutuallyExclusive("radius", "bbox")
	cmd.MarkFlagsMutuallyExclusive("coordinates", "bbox")
}
