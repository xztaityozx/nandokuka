// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//

package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-pipeline"
	"github.com/spf13/cobra"
)

// asciiCmd represents the ascii command
var asciiCmd = &cobra.Command{
	Use:   "ascii",
	Short: "Convert to ASCII Obfuscation",
	Long: `Convert ShellGei to ASCII Obfuscation
	ex)
	date => $'\\x64\\x61\\x74\\x65'`,
	Run: func(cmd *cobra.Command, args []string) {
		err := redirect(args)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		var rt string

		if decodeFlag {
			rt = convertToString()
		} else {
			rt = convertToASCII()
		}

		fmt.Print(rt)
	},
}

func init() {
	rootCmd.AddCommand(asciiCmd)
}

func convertToString() string {
	out, err := pipeline.Output(
		[]string{"cat", TMP_PATH},
		[]string{"xargs", "echo", "-e"},
		[]string{"sed", "s/^\\$//g"},
	)

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func convertToASCII() string {
	out, err := pipeline.Output(
		[]string{"xxd", "-ps", TMP_PATH},
		[]string{"sed", "s/../\\\\x&/g"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("$'%s'\n", strings.Trim(string(out), "\n"))
}
