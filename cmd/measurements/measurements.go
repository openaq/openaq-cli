package measurements

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/openaq/openaq-cli/cmd/flags"
	"github.com/openaq/openaq-cli/cmd/internal"
	"github.com/openaq/openaq-go"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	MeasurementsCmd.AddCommand(listCmd)

	flags.AddLimit(MeasurementsCmd)
	flags.AddPage(MeasurementsCmd)
	flags.AddFormat(MeasurementsCmd)
	flags.AddFromDate(MeasurementsCmd)
	flags.AddToDate(MeasurementsCmd)
	flags.AddPeriodName(MeasurementsCmd)
	flags.AddMini(MeasurementsCmd)
	flags.AddParameters(MeasurementsCmd)
}

func parseFlags(flags *pflag.FlagSet) (*openaq.MeasurementsArgs, error) {
	measurementsArgs := &openaq.MeasurementsArgs{}
	baseArgs := openaq.BaseArgs{}
	limit, err := flags.GetInt64("limit")
	if err != nil {
		return nil, err
	}
	baseArgs.Limit = limit
	page, err := flags.GetInt64("page")
	if err != nil {
		return nil, err
	}
	baseArgs.Page = page
	measurementsArgs.BaseArgs = baseArgs
	periodName, err := flags.GetString("period-name")
	if err != nil {
		return nil, err
	}
	measurementsArgs.PeriodName = periodName
	parametersIDs, err := flags.GetInt64Slice("parameters")
	if err != nil {
		return nil, err
	}
	if len(parametersIDs) > 0 {
		parameters := openaq.Parameters{
			IDs: parametersIDs,
		}
		measurementsArgs.Parameters = &parameters
	}
	from, err := flags.GetString("from")

	if err != nil {
		return nil, err
	}
	var dateFrom time.Time
	if from != "" {
		dateFrom, err = time.Parse("2006-01-02", from)
		if err != nil {
			return nil, err
		}
		measurementsArgs.DatetimeFrom = dateFrom
	}
	if err != nil {
		return nil, err
	}
	to, err := flags.GetString("to")
	if err != nil {
		return nil, err
	}
	var dateTo time.Time
	if to != "" {
		dateTo, err = time.Parse("2006-01-02", to)
		if err != nil {
			return nil, err
		}
		measurementsArgs.DatetimeTo = dateTo
	}
	return measurementsArgs, nil
}

var MeasurementsCmd = &cobra.Command{
	Use:   "measurements",
	Short: "OpenAQ measurements",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := internal.CheckAPIKey()
		if err != nil {
			return err
		}
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list [locationsID]",
	Short: "Get measurements by location ID",
	Long:  `Get measurements by location ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		locationsID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		client, err := internal.SetupClient()
		if err != nil {
			fmt.Println("cannot initialize client")
		}
		ctx := context.Background()
		measurementsArgs, err := parseFlags(cmd.Flags())
		if err != nil {
			panic(err)
		}
		measurements, err := client.GetLocationMeasurements(ctx, locationsID, *measurementsArgs)
		if err != nil {
			panic(err)
		}
		res := internal.FormatResult(measurements, cmd.Flags())
		fmt.Println(res)
	},
}
