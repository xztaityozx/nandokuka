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
	Short: "ASCII難読化します",
	Long: `ASCII難読化シェル芸コンバーター
	Usage : nandokuka ascii [-d|--decode] [FILE]

	[FILE]を空にするとstdinから受け取ります

	ex)
	date => $'\x64\x61\x74\x65'
	
ASCII難読化は以下みたいにワンライナーをまとめて変換します

	seq 30 | awk 'NR%2==0{print}' => $'\x73\x65\x71\x20\x33\x30\x7c\x61\x77\x6b\x20\x27\x4e\x52\x25\x32\x3d\x3d\x30\x7b\x70\x72\x69\x6e\x74\x7d\x27'

変換された出力を eval すれば実行できます
eval しなくても実行できるようにするにはシェル芸をパースする必要があるので大変です`,
	Run: func(cmd *cobra.Command, args []string) {
		err := redirect(args)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		var rt string

		if decodeFlag {
			rt = convertToStringFromASCII()
		} else {
			rt = convertToASCII()
		}

		fmt.Print(rt)
	},
}

func init() {
	rootCmd.AddCommand(asciiCmd)
}

func convertToStringFromASCII() string {
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
