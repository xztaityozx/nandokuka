// Copyright © 2018 xztaityozx
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/mattn/go-pipeline"
	"github.com/spf13/cobra"
)

// gzipCmd represents the gzip command
var gzipCmd = &cobra.Command{
	Use:   "gzip",
	Short: "gzipを使って難読化します",
	Long: `gzip,gunzip難読化コンバーターです
	Usage : nandokuka [-d|--decode] gzip [-e|--excutable] [FILE]
	
	[FILE]を空にするとstdinから受け取ります

	ex)
	date => 1f8b080078d81b5b00034b492c4905007a379eaa04000000

gzipにワンライナーを流し込み、得られた結果をxxdに通してから出力します。xxdのオプションは-psです

-e|--excutableをつけると実行可能な形で出力します

	ex)
	date => eval $(echo -n 1f8b0800ded81b5b00034b492c4905007a379eaa04000000|xxd -ps -r|gunzip)

-d|--decodeが有効のとき、つまりデコードするときには "eval $(echo -n" と "|xxd -ps -r|gunzip)" の部分を外してください


	`,
	Run: func(cmd *cobra.Command, args []string) {
		path := ""
		if len(args) == 1 {
			path = args[1]
		}

		res := readFromFile(path, func(s string) string {
			if decodeFlag {
				return gunzipCommand(s)
			} else {
				return gzipCommand(s)
			}
		})

		fmt.Println(res)
	},
}

func gunzipCommand(s string) string {
	out, err := pipeline.Output(
		[]string{"echo", "-n", s},
		[]string{"xxd", "-ps", "-r"},
		[]string{"gunzip"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return bstring(out)
}

func gzipCommand(s string) string {

	out, err := pipeline.Output(
		[]string{"echo", "-n", s},
		[]string{"gzip", "-fc"},
		[]string{"xxd", "-ps"},
	)

	if err != nil {
		log.Fatal(err)
	}

	res := strings.Trim(bstring(out), "\n")

	if excutableFlag {
		res = "eval $(echo -n " + res + "|xxd -ps -r|gunzip)"
	}

	return res
}

func init() {
	rootCmd.AddCommand(gzipCmd)
	gzipCmd.Flags().BoolVarP(&excutableFlag, "excutable", "e", false, "実行可能な形にして出力します")
}
