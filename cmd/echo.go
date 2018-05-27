// Copyright © 2018 xztaityozx
//

package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"mvdan.cc/sh/syntax"

	"github.com/mattn/go-pipeline"
	"github.com/spf13/cobra"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "一文字一文字echoする難読化をします",
	Long: `一文字一文字echoする難読化コンバーターです
	Usage : nandokuka [-d|--decode] echo [-v|--verbose] [-r|--random] [FILE]

	[FILE]を空にするとstdinから受け取ります

	ex)
	date => $(echo d)$(echo a)$(echo t)$(echo e)

一文字ずつechoして難読化します
-rをつけるとechoの他にprintfも使うようになります

この処理の途中でシングルクォートが削除されることに注意してください
これはawkやsedを正確にエンコードするためにしています
もう少しマシなパースができるようになったらがんばります
`,
	Run: func(cmd *cobra.Command, args []string) {
		var file *os.File
		if len(args) == 0 {
			file = os.Stdin
		} else {
			file, _ = os.Open(args[0])
		}

		defer file.Close()
		s := bufio.NewScanner(file)

		parser := syntax.NewParser(syntax.Variant(syntax.LangBash))
		printer := syntax.NewPrinter()

		var rt string
		for s.Scan() {
			line := s.Text()

			prog, _ := parser.Parse(strings.NewReader(line), "")
			syntax.Simplify(prog)

			var writeBuf bytes.Buffer
			writeBuf.Reset()
			printer.Print(&writeBuf, prog)
			res := string(writeBuf.Bytes())

			if verbFlag {
				fmt.Printf("shfmt : %s\n", res)
			}

			if decodeFlag {
				rt += convertToStringFromEcho(res)
			} else {
				get := strings.Replace(res, "'", "", -1)
				rt += convertToEcho(get)
			}
		}

		fmt.Printf("%s", rt)
	},
}

func convertToStringFromEcho(s string) string {
	out, _ := pipeline.Output(
		[]string{"echo", "echo", "-n", "\"", s, "\""},
		[]string{"bash"},
	)
	return string(out)
}

func convertToEcho(s string) string {
	var rt string
	for i := 0; i < len(s); i++ {
		x := rand.Int31() % 2
		y := rand.Int31() % 2
		f := !(x == 0 && randomFlag)
		node := strings.Replace(strconv.Quote(string(s[i])), "`", "\\`", -1)
		var item string

		if ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z') {
			if f {
				//echo
				item = surround(y == 0, "echo "+node)
			} else {
				item = surround(y == 0, "printf "+node+"")
			}

		} else if s[i] == ' ' {
			item = " "
		} else {
			item = string(s[i])
		}

		if verbFlag {
			fmt.Printf("%s => %s\n", string(s[i]), item)
		}

		rt += item
	}
	if verbFlag {
		fmt.Printf("\n")
	}
	return rt
}

var randomFlag bool
var verbFlag bool

func init() {
	rootCmd.AddCommand(echoCmd)

	echoCmd.Flags().BoolVarP(&randomFlag, "random", "r", false, "echoとprintfを混ぜます")
	echoCmd.Flags().BoolVarP(&verbFlag, "verbose", "v", false, "変換の過程も出力します")
}
