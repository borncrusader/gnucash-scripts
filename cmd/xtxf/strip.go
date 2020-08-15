package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

type StripCommand struct {
	Command *flag.FlagSet

	RemovePII   bool
	RemoveMoney bool
	InputFile   string
	OutputFile  string
}

func handleStripCommand(command StripCommand) error {
	if command.InputFile == "" || command.OutputFile == "" {
		return errReqdArgMissing
	}

	f, err := os.Open(command.InputFile)
	if err != nil {
		return err
	}

	c := csv.NewReader(f)

	records, err := c.ReadAll()
	if err != nil {
		return err
	}

	//if command.RemovePII {

	//}

	t := table.NewWriter()

	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Number: 4,
			Colors: text.Colors{text.BgBlack, text.FgRed},
		},
		{
			Number: 5,
			Colors: text.Colors{text.BgBlack, text.FgRed},
		},
		{
			Number: 6,
			Colors: text.Colors{text.BgBlack, text.FgMagenta},
		},
	})

	for _, record := range records {
		row := make([]interface{}, len(record))
		for j, entry := range record {
			row[j] = entry
		}

		t.AppendRow(row)
	}

	fmt.Println(t.Render())

	return nil
}

func isHeaderRow(row []string) float64 {
	ret := 0.0

	for _, column := range row {
		if strings.EqualFold(column, "Description") {
			ret += 1
		} else if strings.EqualFold(column, "Debit") {
			ret += 1
		} else if strings.EqualFold(column, "Credit") {
			ret += 1
		} else if strings.EqualFold(column, "Date") {
			ret += 1
		}
	}

	return ret / float64(len(row))
}
