package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	commandName    = "pn-checker"
	diligentOption = "diligent"
	abuse          = `
   _____ _     _       _       __
  |_   _| |__ (_)_ __ | | __  / _| ___  _ __
    | | | '_ \| | '_ \| |/ / | |_ / _ \| '__|
    | | | | | | | | | |   <  |  _| (_) | |
    |_| |_| |_|_|_| |_|_|\_\ |_|  \___/|_|
                                   _  __        _               _     _   _
   _   _  ___  _   _ _ __ ___  ___| |/ _|   ___| |_ _   _ _ __ (_) __| | | |
  | | | |/ _ \| | | | '__/ __|/ _ \ | |_   / __| __| | | | '_ \| |/ _  | | |
  | |_| | (_) | |_| | |  \__ \  __/ |  _|  \__ \ |_| |_| | |_) | | (_| | |_|
   \__, |\___/ \__,_|_|  |___/\___|_|_|( ) |___/\__|\__,_| .__/|_|\__,_| (_)
   |___/                               |/                |_|
 `
)

var (
	version            = "v1.1.0"
	primeNumberChecker = &cobra.Command{
		Use:     commandName,
		Short:   fmt.Sprintf("%s is a slightly silly and useless command line tool.", commandName),
		Long:    fmt.Sprintf("%s checks whether the number specified is a prime number or not.\nAnd sometimes, it says \"Think for yourself, stupid!\" to you.", commandName),
		Example: fmt.Sprintf("%s 997", commandName),
		Version: version,
		Args:    cobra.RangeArgs(1, 1),
		RunE:    RunCommand,
	}
	isDiligent bool
)

// Execute runs command.
// If error is occurred, exit with status code 1.
func Execute() {
	if err := primeNumberChecker.Execute(); err != nil {
		os.Exit(1)
	}
}

// RunCommand runs command, but returns an error.
func RunCommand(cmd *cobra.Command, args []string) error {
	targetNumber, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil || targetNumber <= 0 {
		return fmt.Errorf("argument must be positive integer, received \"%s\"", args[0])
	}

	// If diligent flag is false, returns harsh words with a certain probability.
	// If it is true, returns the result of checking if the entered value is a prime number.
	if !isDiligent {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(100) >= 50 {
			c := color.New(color.FgHiRed, color.Bold)
			c.Println(abuse)
			return nil
		}
	}
	fmt.Println(checkAndGenerateMessage(targetNumber))

	return nil
}

func checkAndGenerateMessage(target uint64) string {
	switch {
	case target == 0:
		return "The number 0 is not a prime number.\nNot even a natural number."
	case target == 1:
		return "The number 1 is not a prime number."
	default:
		var notPrimeNumber bool
		var i uint64
		for i = 2; i < target; i++ {
			if target%i == 0 {
				notPrimeNumber = true
				break
			}
		}

		if notPrimeNumber {
			return fmt.Sprintf("The number %d is not a prime number.", target)
		}
		return fmt.Sprintf("Congrats!! The number %d is a prime number.", target)
	}
}

func init() {
	// Set customized flags
	primeNumberChecker.PersistentFlags().BoolVarP(&isDiligent, diligentOption, "d", false, "seriously check if it's a prime number")

	// Set customized usage template
	primeNumberChecker.SetUsageTemplate(`Usage:{{if .Runnable}}
  {{.UseLine}} [arg(positive integer)] {{end}}{{if .HasExample}}

Examples:
  {{.Example}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
`)
}
