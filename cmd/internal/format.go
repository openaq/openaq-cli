package internal

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nwidger/jsoncolor"
	"github.com/spf13/pflag"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/openaq/openaq-cli/cmd/config"
	"github.com/openaq/openaq-go"
)

var countriesHeaders = []string{"countries_id", "iso_code", "name", "datetime_first", "datetime_last", "parameters", "locations_count", "measurements_count", "providers_count"}
var countriesMiniHeaders = []string{"countries_id", "iso_code", "name"}
var locationsHeaders = []string{"locations_id", "name", "countries_id", "country_iso", "country_name", "latitude", "longitude"}
var measurementsParametersHeaders = []string{"parameter_id", "parameter_name", "parameter_units", "parameter_display_name"}
var measurementsPeriodHeaders = []string{"periodLabel", "periodInterval", "periodDatetimeFromUTC", "periodDatetimeFromLocal", "periodDatetimeToUTC", "periodDatetimeToLocal"}
var measurementsCoverageHeaders = []string{"expectedCount", "expectedInterval", "observedCount", "observedInterval", "percentComplete", "percentCoverage", "datetimeFromUTC", "datetimeFromLocal", "datetimeToUTC", "datetimeToLocal"}
var measurementsSummaryHeaders = []string{"min", "q02", "q25", "median", "q75", "q98", "max", "sd"}
var measurementsHeaders = appendMany([][]string{measurementsParametersHeaders, {"value"}, measurementsPeriodHeaders, measurementsCoverageHeaders, measurementsSummaryHeaders})
var miniMeasurementsHeaders = []string{"parameter", "datetime_local", "datetime_utc", "period", "value"}
var ownersHeaders = []string{"id", "name", "locationsCount"}
var parametersHeaders = []string{"parameters_id", "name", "display_name", "units", "description", "locations_count", "measurements_count"}
var providersHeaders = []string{"providers_id", "name", "source_name", "export_prefix", "license", "datetime_added", "datetime_first", "datetime_last", "locations_count", "measurements_count", "countries_count"}

func FormatResult(v interface{}, flags *pflag.FlagSet) string {
	jsonFlag, _ := flags.GetBool("json")
	csvFlag, _ := flags.GetBool("csv")
	pretty, _ := flags.GetBool("pretty")
	mini, _ := flags.GetBool("mini")
	if csvFlag || (config.CSVConfig && !jsonFlag) {
		var csvOut string
		switch v := v.(type) {
		case *openaq.LocationsResponse:
			csvOut = writeLocationsCSV(v, locationsHeaders)
		case *openaq.CountriesResponse:
			if mini {
				csvOut = writeMiniCountriesCSV(v, countriesMiniHeaders)
			} else {
				csvOut = writeCountriesCSV(v, countriesHeaders)
			}
		case *openaq.MeasurementsResponse:
			if mini {
				csvOut = writeMiniMeasurementsCSV(v, miniMeasurementsHeaders)
			} else {
				csvOut = writeMeasurementsCSV(v, measurementsHeaders)
			}

		case *openaq.ParametersResponse:
			csvOut = writeParametersCSV(v, parametersHeaders)
		case *openaq.ProvidersResponse:
			csvOut = writeProvidersCSV(v, providersHeaders)

		case *openaq.OwnersResponse:
			csvOut = writeOwnersCSV(v, ownersHeaders)

		default:
			fmt.Println("cannot find type")
		}

		return csvOut
	}
	if jsonFlag || (config.JSONConfig && !csvFlag) {
		var out []byte
		var err error
		if pretty || config.PrettyConfig {
			out, err = jsoncolor.MarshalIndent(v, "", "   ")
		} else {
			out, err = json.Marshal(v)
		}
		if err != nil {
			panic(err)
		}
		return string(out)
	}
	switch v := v.(type) {
	case *openaq.LocationsResponse:
		return writeLocationsTable(v, locationsHeaders)
	case *openaq.MeasurementsResponse:
		if mini {
			return writeMiniMeasurementsTable(v, miniMeasurementsHeaders)
		} else {
			return writeMeasurementsTable(v, measurementsHeaders)
		}
	case *openaq.CountriesResponse:
		if mini {
			return writeMiniCountriesTable(v, countriesMiniHeaders)
		} else {
			return writeCountriesTable(v, countriesHeaders)
		}
	case *openaq.OwnersResponse:
		return writeOwnersTable(v, ownersHeaders)

	case *openaq.ParametersResponse:
		return writeParametersTable(v, parametersHeaders)
	case *openaq.ProvidersResponse:
		return writeProvidersTable(v, providersHeaders)
	default:
		return ""
	}
}

// utility func for writing pretty table header
func writeTableHeader(tw table.Writer, headerVals []string) {
	header := make(table.Row, 0, len(headerVals))
	// convert from []string to []interface{} for table.Row
	for _, c := range headerVals {
		header = append(header, c)
	}
	tw.AppendHeader(header)
}

// countries
func writeCountriesCSV(countries *openaq.CountriesResponse, headers []string) string {

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.Write(headers)

	for _, s := range countries.Results {
		var record []string
		record = append(record, strconv.FormatInt(s.ID, 10))
		record = append(record, s.Code)
		record = append(record, s.Name)
		record = append(record, s.DatetimeFirst.Format(time.RFC3339))
		record = append(record, s.DatetimeLast.Format(time.RFC3339))
		record = append(record, "")
		record = append(record, strconv.FormatInt(s.LocationsCount, 10))
		record = append(record, strconv.FormatInt(s.MeasurementsCount, 10))
		record = append(record, strconv.FormatInt(s.ProvidersCount, 10))
		w.Write(record)

	}
	w.Flush()
	return buf.String()
}

func writeMiniCountriesCSV(countries *openaq.CountriesResponse, headers []string) string {

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.Write(headers)

	for _, s := range countries.Results {
		var record []string
		record = append(record, strconv.FormatInt(s.ID, 10))
		record = append(record, s.Code)
		record = append(record, s.Name)
		record = append(record, "")
		w.Write(record)

	}
	w.Flush()
	return buf.String()
}

func writeCountriesTable(countries *openaq.CountriesResponse, headers []string) string {
	var columns = len(headers)
	tw := table.NewWriter()
	writeTableHeader(tw, headers)
	for _, s := range countries.Results {
		row := make(table.Row, 0, columns)
		row = append(row, strconv.FormatInt(s.ID, 10))
		row = append(row, s.Code)
		row = append(row, s.Name)
		row = append(row, s.DatetimeFirst.Format(time.RFC3339))
		row = append(row, s.DatetimeLast.Format(time.RFC3339))
		row = append(row, joinParamDisplayNames(s.Parameters))
		row = append(row, strconv.FormatInt(s.LocationsCount, 10))
		row = append(row, strconv.FormatInt(s.MeasurementsCount, 10))
		row = append(row, strconv.FormatInt(s.ProvidersCount, 10))
		tw.AppendRow(row)
	}
	return tw.Render()
}

func joinParamDisplayNames(params []openaq.ParameterBase) string {
	var builder strings.Builder
	for i, param := range params {
		builder.WriteString(param.DisplayName)
		if i < len(params)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}

func writeMiniCountriesTable(countries *openaq.CountriesResponse, headers []string) string {
	var columns = len(headers)
	tw := table.NewWriter()
	writeTableHeader(tw, headers)
	for _, s := range countries.Results {
		row := make(table.Row, 0, columns)
		row = append(row, strconv.FormatInt(s.ID, 10))
		row = append(row, s.Code)
		row = append(row, s.Name)
		tw.AppendRow(row)
	}
	return tw.Render()
}

// owners
func writeOwnersTable(owners *openaq.OwnersResponse, headers []string) string {
	var columns = len(headers)
	tw := table.NewWriter()
	writeTableHeader(tw, headers)
	for _, s := range owners.Results {
		row := make(table.Row, 0, columns)
		row = append(row, strconv.FormatInt(s.ID, 10))
		row = append(row, s.Name)
		row = append(row, strconv.FormatInt(s.LocationsCount, 10))
		tw.AppendRow(row)
	}
	return tw.Render()
}

func writeOwnersCSV(data *openaq.OwnersResponse, headers []string) string {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	writer.Write(headers)

	for _, s := range data.Results {
		writer.Write([]string{
			strconv.FormatInt(s.ID, 10),
			s.Name,
			strconv.FormatInt(s.LocationsCount, 10),
		})
	}

	writer.Flush()
	return buffer.String()
}

// parameters
func writeParametersCSV(parameters *openaq.ParametersResponse, headers []string) string {

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.Write(headers)

	for _, s := range parameters.Results {
		var record []string
		record = append(record, strconv.FormatInt(s.ID, 10))
		record = append(record, s.Name)
		record = append(record, s.DisplayName)
		record = append(record, s.Units)
		record = append(record, s.Description)
		record = append(record, strconv.FormatInt(s.LocationsCount, 10))
		record = append(record, strconv.FormatInt(s.MeasurementsCount, 10))
		w.Write(record)

	}
	w.Flush()
	return buf.String()
}

func writeProvidersTable(countries *openaq.ProvidersResponse, headers []string) string {
	var columns = len(headers)
	tw := table.NewWriter()
	writeTableHeader(tw, headers)
	for _, s := range countries.Results {
		row := make(table.Row, 0, columns)
		row = append(row, strconv.FormatInt(s.ID, 10))
		row = append(row, s.Name)
		row = append(row, s.SourceName)
		row = append(row, s.ExportPrefix)
		row = append(row, s.License)
		row = append(row, s.DatetimeAdded.Format(time.RFC3339))
		row = append(row, s.DatetimeFirst.Format(time.RFC3339))
		row = append(row, s.DatetimeLast.Format(time.RFC3339))
		row = append(row, strconv.FormatInt(s.LocationsCount, 10))
		row = append(row, strconv.FormatInt(s.MeasurementsCount, 10))
		row = append(row, strconv.FormatInt(s.CountriesCount, 10))
		tw.AppendRow(row)
	}
	return tw.Render()

}

func writeProvidersCSV(providers *openaq.ProvidersResponse, headers []string) string {

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.Write(headers)

	for _, s := range providers.Results {
		var record []string
		record = append(record, strconv.FormatInt(s.ID, 10))
		record = append(record, s.Name)
		record = append(record, s.SourceName)
		record = append(record, s.ExportPrefix)
		record = append(record, s.License)
		record = append(record, s.DatetimeAdded.Format(time.RFC3339))
		record = append(record, s.DatetimeFirst.Format(time.RFC3339))
		record = append(record, s.DatetimeLast.Format(time.RFC3339))
		record = append(record, strconv.FormatInt(s.LocationsCount, 10))
		record = append(record, strconv.FormatInt(s.MeasurementsCount, 10))
		record = append(record, strconv.FormatInt(s.CountriesCount, 10))
		w.Write(record)
	}
	w.Flush()
	return buf.String()
}

func writeParametersTable(countries *openaq.ParametersResponse, headers []string) string {
	var columns = len(headers)
	tw := table.NewWriter()
	writeTableHeader(tw, headers)
	for _, s := range countries.Results {
		row := make(table.Row, 0, columns)
		row = append(row, strconv.FormatInt(s.ID, 10))
		row = append(row, s.Name)
		row = append(row, s.DisplayName)
		row = append(row, s.Units)
		row = append(row, s.Description)
		row = append(row, strconv.FormatInt(s.LocationsCount, 10))
		row = append(row, strconv.FormatInt(s.MeasurementsCount, 10))
		tw.AppendRow(row)
	}
	return tw.Render()

}

// locations
func writeLocationsCSV(locations *openaq.LocationsResponse, headers []string) string {

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.Write(headers)

	for _, s := range locations.Results {
		var record []string
		record = append(record, strconv.FormatInt(s.ID, 10))
		record = append(record, s.Name)
		record = append(record, strconv.FormatInt(s.Country.ID, 10))
		record = append(record, string(s.Country.Code))
		record = append(record, string(s.Country.Name))
		record = append(record, fmt.Sprintf("%f", s.Coordinates.Latitude))
		record = append(record, fmt.Sprintf("%f", s.Coordinates.Longitude))
		w.Write(record)

	}
	w.Flush()
	return buf.String()
}

func writeLocationsTable(locations *openaq.LocationsResponse, headers []string) string {
	var columns = len(headers)
	tw := table.NewWriter()
	writeTableHeader(tw, headers)
	for _, s := range locations.Results {
		row := make(table.Row, 0, columns)
		row = append(row, strconv.FormatInt(s.ID, 10))
		row = append(row, s.Name)
		row = append(row, strconv.FormatInt(s.Country.ID, 10))
		row = append(row, string(s.Country.Code))
		row = append(row, string(s.Country.Name))
		row = append(row, fmt.Sprintf("%f", s.Coordinates.Latitude))
		row = append(row, fmt.Sprintf("%f", s.Coordinates.Longitude))
		tw.AppendRow(row)
	}
	return tw.Render()

}

// measurements
func writeMeasurementsCSV(measurements *openaq.MeasurementsResponse, headers []string) string {

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.Write(headers)
	for _, s := range measurements.Results {
		var record []string
		record = append(record, strconv.FormatInt(s.Parameter.ID, 10))
		record = append(record, s.Parameter.Name)
		record = append(record, s.Parameter.Units)
		record = append(record, s.Parameter.DisplayName)
		record = append(record, strconv.FormatFloat(s.Value, 'f', -1, 64))
		record = append(record, s.Period.Interval)
		record = append(record, s.Period.DatetimeFrom.UTC.Format(time.RFC3339))
		record = append(record, s.Period.DatetimeFrom.Local.Format(time.RFC3339))
		record = append(record, s.Period.DatetimeTo.UTC.Format(time.RFC3339))
		record = append(record, s.Period.DatetimeTo.Local.Format(time.RFC3339))
		record = append(record, strconv.FormatInt(s.Coverage.ExpectedCount, 10))
		record = append(record, s.Coverage.ExpectedInterval)
		record = append(record, strconv.FormatInt(s.Coverage.ObservedCount, 10))
		record = append(record, s.Coverage.ObservedInterval)
		record = append(record, strconv.FormatFloat(s.Coverage.PercentComplete, 'f', -1, 64))
		record = append(record, strconv.FormatFloat(s.Coverage.PercentCoverage, 'f', -1, 64))
		record = append(record, s.Coverage.DatetimeFrom.UTC.Format(time.RFC3339))
		record = append(record, s.Coverage.DatetimeFrom.Local.Format(time.RFC3339))
		record = append(record, s.Coverage.DatetimeTo.UTC.Format(time.RFC3339))
		record = append(record, s.Coverage.DatetimeTo.Local.Format(time.RFC3339))
		record = append(record, strconv.FormatFloat(s.Summary.Min, 'f', -1, 64))
		record = append(record, strconv.FormatFloat(s.Summary.Q02, 'f', -1, 64))
		record = append(record, strconv.FormatFloat(s.Summary.Q25, 'f', -1, 64))
		record = append(record, strconv.FormatFloat(s.Summary.Median, 'f', -1, 64))
		record = append(record, strconv.FormatFloat(s.Summary.Q75, 'f', -1, 64))
		record = append(record, strconv.FormatFloat(s.Summary.Q98, 'f', -1, 64))
		record = append(record, strconv.FormatFloat(s.Summary.Max, 'f', -1, 64))
		record = append(record, strconv.FormatFloat(s.Summary.StdDev, 'f', -1, 64))
		w.Write(record)

	}
	w.Flush()
	return buf.String()
}

func writeMeasurementsTable(measurements *openaq.MeasurementsResponse, headers []string) string {
	fmt.Println(headers)
	var columns = len(headers)
	tw := table.NewWriter()
	writeTableHeader(tw, headers)
	for _, s := range measurements.Results {
		row := make(table.Row, 0, columns)
		row = append(row, strconv.FormatInt(s.Parameter.ID, 10))
		row = append(row, s.Parameter.Name)
		row = append(row, s.Parameter.Units)
		row = append(row, s.Parameter.DisplayName)
		row = append(row, strconv.FormatFloat(s.Value, 'f', -1, 64))
		row = append(row, s.Period.Interval)
		row = append(row, s.Period.DatetimeFrom.UTC.Format(time.RFC3339))
		row = append(row, s.Period.DatetimeFrom.Local.Format(time.RFC3339))
		row = append(row, s.Period.DatetimeTo.UTC.Format(time.RFC3339))
		row = append(row, s.Period.DatetimeTo.Local.Format(time.RFC3339))
		row = append(row, strconv.FormatInt(s.Coverage.ExpectedCount, 10))
		row = append(row, s.Coverage.ExpectedInterval)
		row = append(row, strconv.FormatInt(s.Coverage.ObservedCount, 10))
		row = append(row, s.Coverage.ObservedInterval)
		row = append(row, strconv.FormatFloat(s.Coverage.PercentComplete, 'f', -1, 64))
		row = append(row, strconv.FormatFloat(s.Coverage.PercentCoverage, 'f', -1, 64))
		row = append(row, s.Coverage.DatetimeFrom.UTC.Format(time.RFC3339))
		row = append(row, s.Coverage.DatetimeFrom.Local.Format(time.RFC3339))
		row = append(row, s.Coverage.DatetimeTo.UTC.Format(time.RFC3339))
		row = append(row, s.Coverage.DatetimeTo.Local.Format(time.RFC3339))
		row = append(row, strconv.FormatFloat(s.Summary.Min, 'f', -1, 64))
		row = append(row, strconv.FormatFloat(s.Summary.Q02, 'f', -1, 64))
		row = append(row, strconv.FormatFloat(s.Summary.Q25, 'f', -1, 64))
		row = append(row, strconv.FormatFloat(s.Summary.Median, 'f', -1, 64))
		row = append(row, strconv.FormatFloat(s.Summary.Q75, 'f', -1, 64))
		row = append(row, strconv.FormatFloat(s.Summary.Q98, 'f', -1, 64))
		row = append(row, strconv.FormatFloat(s.Summary.Max, 'f', -1, 64))
		row = append(row, strconv.FormatFloat(s.Summary.StdDev, 'f', -1, 64))
		tw.AppendRow(row)
	}
	return tw.Render()

}

func writeMiniMeasurementsCSV(measurements *openaq.MeasurementsResponse, headers []string) string {
	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.Write(headers)
	for _, s := range measurements.Results {
		var record []string
		record = append(record, fmt.Sprintf("%s %s", s.Parameter.Name, s.Parameter.Units))
		record = append(record, s.Period.DatetimeFrom.Local.Format(time.RFC3339))
		record = append(record, s.Period.DatetimeTo.UTC.Format(time.RFC3339))
		record = append(record, s.Period.Interval)
		record = append(record, strconv.FormatFloat(s.Value, 'f', -1, 64))
		w.Write(record)
	}
	w.Flush()
	return buf.String()
}

// {"parameter", "datetime_local", "datetime_utc", "period", "value"}
func writeMiniMeasurementsTable(measurements *openaq.MeasurementsResponse, headers []string) string {
	var columns = len(headers)
	tw := table.NewWriter()
	writeTableHeader(tw, headers)
	for _, s := range measurements.Results {
		row := make(table.Row, 0, columns)
		row = append(row, fmt.Sprintf("%s %s", s.Parameter.Name, s.Parameter.Units))
		row = append(row, s.Period.DatetimeFrom.Local.Format(time.RFC3339))
		row = append(row, s.Period.DatetimeTo.UTC.Format(time.RFC3339))
		row = append(row, s.Period.Interval)
		row = append(row, s.Value)
		tw.AppendRow(row)
	}
	return tw.Render()

}
