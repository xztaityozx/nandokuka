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
	"strings"

	"github.com/spf13/cobra"
)

// symbolCmd represents the symbol command
var symbolCmd = &cobra.Command{
	Use:   "symbol",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("symbol called")
	},
}

var numTable = map[string]string{
	"0": "1-1",
	"1": "1",
	"2": "2",
	"3": "1+2",
	"4": "2+2",
	"5": "2*2+1",
	"6": "2+2*2",
	"7": "2*2*2-1",
	"8": "2*2+2*2",
	"9": "2*2*2+1",
}

func encodeToSymbolOnly(s string) string {
	var res string

	var convTable = map[string]int{
		"z":  1,
		"y":  2,
		"x":  3,
		"w":  4,
		"v":  5,
		"u":  6,
		"t":  7,
		"s":  8,
		"r":  9,
		"q":  10,
		"p":  11,
		"o":  12,
		"n":  13,
		"m":  14,
		"l":  15,
		"k":  16,
		"j":  17,
		"i":  18,
		"h":  19,
		"g":  20,
		"f":  21,
		"e":  22,
		"d":  23,
		"c":  24,
		"b":  25,
		"a":  26,
		"`":  27,
		"_":  28,
		"^":  29,
		"]":  30,
		"\\": 31,
		"[":  32,
		"Z":  33,
		"Y":  34,
		"X":  35,
		"W":  36,
		"V":  37,
		"U":  38,
		"T":  39,
		"S":  40,
		"R":  41,
		"Q":  42,
		"P":  43,
		"O":  44,
		"N":  45,
		"M":  46,
		"L":  47,
		"K":  48,
		"J":  49,
		"I":  50,
		"H":  51,
		"G":  52,
		"F":  53,
		"E":  54,
		"D":  55,
		"C":  56,
		"B":  57,
		"A":  58,
	}

	if prefixFlag {
		res = "A=$(. 2>&1);A=${A##*.};${A:1$((2*2*2+1)):1}${A:$((2+2)):1}${A:1$((2*2+2*2)):1} -- {z..A};"
	}

	for key, value := range convTable {
		strings.Replace(s, key, fmt.Sprintf("${A:%d:1}", value), -1)
	}

	for key, value := range numTable {
		strings.Replace(s, key, value, -1)
	}
	return res + s
}

var prefixFlag bool
var superFlag bool

func init() {
	rootCmd.AddCommand(symbolCmd)
	symbolCmd.Flags().BoolVarP(&superFlag, "super", "s", false, "超記号オンリー難読化します")
	symbolCmd.Flags().BoolVarP(&prefixFlag, "prefix", "p", false, "記号オンリー難読化に必要な材料を一緒に出力します")
}
