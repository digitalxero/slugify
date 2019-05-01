package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/digitalxero/slugify"
)

type slugifierFunc = func(text string, maxLen int) string

var (
	pkgName    = "slugify"
	version    = "0.0.1"
	repoBranch = "develop"
	cmdRoot    = &cobra.Command{
		Use:   pkgName,
		Short: "CLI Tool to slugify a string",
	}
	lowerOnly = false
	maxLen         = 0
	ok = slugify.OK
	dash = slugify.TO_DASH
	slugifier slugifierFunc
)

func main() {
	cmdRoot.Flags().IntVarP(
		&maxLen,
		"max-len",
		"l",
		maxLen,
		``)

	cmdRoot.Flags().BoolVarP(
		&lowerOnly,
		"lower",
		"",
		lowerOnly,
		``)

	cmdRoot.Flags().StringVarP(
		&ok,
		"ok",
		"",
		ok,
		`Non alphanumeric values that are OK to have in your output`)

	cmdRoot.Flags().StringVarP(
		&dash,
		"to-dash",
		"",
		dash,
		`Convert these to a dash instead of stripping them from the output`)

	cmdRoot.Run = run
	cmdRoot.PersistentPreRun = preReun

	if err := startCLI(); err != nil {
		os.Exit(127)
	}
}

func startCLI() (err error) {
	// Run!
	return cmdRoot.Execute()
}

func preReun(c *cobra.Command, args []string) {
	slugifier = slugify.IDify
	if lowerOnly {
		slugifier = slugify.Slugify
	}
	slugify.OK = ok
	slugify.TO_DASH = dash
}

func run(c *cobra.Command, args []string) {
	data := strings.Join(args, " ")
	data = slugifier(data, maxLen)

	fmt.Println(data)
}