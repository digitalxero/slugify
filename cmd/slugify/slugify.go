package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/digitalxero/slugify"
)

var (
	pkgName    = "slugify"
	version    = "0.0.1"
	repoBranch = "develop"
	cmdRoot    = &cobra.Command{
		Use:   pkgName,
		Short: "CLI Tool to slugify a string",
	}
	maxLen         = 0
	ok = slugify.ID_OK
	dash = slugify.ID_TO_DASH
)

func main() {
	cmdRoot.PersistentFlags().IntVarP(
		&maxLen,
		"max-len",
		"l",
		maxLen,
		``)

	cmdRoot.PersistentFlags().StringVarP(
		&ok,
		"ok",
		"",
		ok,
		`Non alphanumeric values that are OK to have in your output`)

	cmdRoot.PersistentFlags().StringVarP(
		&dash,
		"to-dash",
		"",
		dash,
		`Convert these to a dash instead of stripping them from the output`)

	cmdRoot.Run = run

	if err := startCLI(); err != nil {
		os.Exit(127)
	}
}

func startCLI() (err error) {
	// Run!
	return cmdRoot.Execute()
}

func run(c *cobra.Command, args []string) {
	data := strings.Join(args, " ")
	slugify.ID_OK = ok
	slugify.ID_TO_DASH = dash
	data = slugify.IDify(data)

	if maxLen > 0 && len(data) > maxLen {
		data = slugify.IDify(data[0:maxLen])
	}

	fmt.Println(data)
}