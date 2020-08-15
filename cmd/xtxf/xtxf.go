package main

import (
	"errors"
	"flag"
	"os"
)

var (
	errReqdArgMissing error = errors.New("required argument missing")
)

func main() {
	var stripCommand StripCommand

	stripCommand.Command = flag.NewFlagSet("strip", flag.ExitOnError)

	// strip subcommand
	stripCommand.Command.BoolVar(&stripCommand.RemovePII, "remove-pii", true, "Remove PII columns")
	stripCommand.Command.BoolVar(&stripCommand.RemoveMoney, "remove-money", true, "Remove monetary columns")
	stripCommand.Command.StringVar(&stripCommand.InputFile, "input", "", "Input file (required")
	stripCommand.Command.StringVar(&stripCommand.OutputFile, "output", "", "Output file (required")

	switch os.Args[1] {
	case "strip":
		stripCommand.Command.Parse(os.Args[2:])
		handleStripCommand(stripCommand)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
