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
		path := ""
		if len(args) == 1 {
			path = args[0]
		}

		res := readFromFile(path, func(s string) string {
			if decodeFlag {
				return ""
			} else {
				if superFlag {
					return encodeToSuperSymbolOnly(s)
				} else {
					return encodeToSymbolOnly(s)
				}
			}
		})
		fmt.Println(strings.Trim(res, "'"))
	},
}

var numTable = map[string]string{
	"0": "$((1-1))",
	"1": "1",
	"2": "2",
	"3": "$((1+2))",
	"4": "$((2+2))",
	"5": "$((2*2+1))",
	"6": "$((2+2*2))",
	"7": "$((2*2*2-1))",
	"8": "$((2*2+2*2))",
	"9": "$((2*2*2+1))",
}

var prefixCommands = map[bool]string{
	true:  `__=$(. 2>&1);__=${__##*.};___=$(${__:$((2*2*2-1)):1}${__:1$((2*2+2*2)):1} 2>&1);___=${___##*]};____="$(${___:1$((1+2)):1}${___:2$((2*2*2-1)):1} ${___:$((2*2*2+1)):4}|${___:2$((2*2*2+1)):1}${___:2:1}${___:12:1}${___:1$((2+2)):1} -${___:1$((2*2*2-1)):1} .|${___:2$((2*2*2-1)):1}${___:1$((2*2*2-1)):2}${___:$((2*2*2-1)):1} -${___:2:1}${___:2$((2+2*2)):1}|${___:$((2*2*2-1)):1}${___:2$((2*2+2*2)):1}${__:2:2} -$((2*2+2*2))0)";____=${____%/*};${___:12:1}${____:$((2*2+2*2)):1}${___:2$((2*2+2*2)):1}${___:1$((1+2)):1} "${___:2$((2*2*2-1)):1}${___:12:1}${___:$((2*2*2-1)):1} -- {${____:$((1-1)):1}..${____:1$((1-1))$((2+2*2)):1}}";`,
	false: "A=$(. 2>&1);A=${A##*.};${A:1$((2*2*2+1)):1}${A:$((2+2)):1}${A:1$((2*2+2*2)):1} -- {z..A};",
}

func encodeToSuperSymbolOnly(s string) string {
	var convTable = map[string]string{
		"z": "____:0:1",
		"y": "___:3:1",
		"x": "____:4:1",
		"w": "____:6:1",
		"v": "____:8:1",
		"u": "__:14:1",
		"t": "__:18:1",
		"s": "__:19:1",
		"r": "__:12:1",
		"q": "____:18:1",
		"p": "___:14:1",
		"o": "___:17:1",
		"n": "__:5:1",
		"m": "__:7:1",
		"l": "__:3:1",
		"k": "____:30:1",
		"j": "____:32:1",
		"i": "__:2:1",
		"h": "___:11:1",
		"g": "__:13:1",
		"f": "__:1:1",
		"e": "__:4:1",
		"d": "____:44:1",
		"c": "____:46:1",
		"b": "____:48:1",
		"a": "__:6:1",
		"Z": "____:60:1",
		"Y": "____:62:1",
		"X": "____:64:1",
		"W": "____:66:1",
		"U": "____:68:1",
		"T": "___:1:1",
		"S": "____:72:1",
		"R": "____:74:1",
		"Q": "____:76:1",
		"P": "____:78:1",
		"O": "____:80:1",
		"N": "____:82:1",
		"M": "____:84:1",
		"L": "____:86:1",
		"K": "____:88:1",
		"J": "@:95:1",
		"I": "____:90:1",
		"H": "____:92:1",
		"G": "____:94:1",
		"F": "____:96:1",
		"E": "____:98:1",
		"D": "____:100:1",
		"C": "____:102:1",
		"B": "____:104:1",
		"A": "____:106:1",
	}

	var res string
	if prefixFlag {
		res = prefixCommands[superFlag]
	}

	for key, value := range convTable {
		s = strings.Replace(s, key, fmt.Sprintf("${%s}", value), -1)
	}
	for key, value := range numTable {
		s = strings.Replace(s, key, fmt.Sprintf("%s", value), -1)
	}

	return res + s
}

func encodeToSymbolOnly(s string) string {
	var res string

	var convTable = map[string]int{
		"z": 1,
		"y": 2,
		"x": 3,
		"w": 4,
		"v": 5,
		"u": 6,
		"t": 7,
		"s": 8,
		"r": 9,
		"q": 10,
		"p": 11,
		"o": 12,
		"n": 13,
		"m": 14,
		"l": 15,
		"k": 16,
		"j": 17,
		"i": 18,
		"h": 19,
		"g": 20,
		"f": 21,
		"e": 22,
		"d": 23,
		"c": 24,
		"b": 25,
		"a": 26,
		"Z": 33,
		"Y": 34,
		"X": 35,
		"W": 36,
		"V": 37,
		"U": 38,
		"T": 39,
		"S": 40,
		"R": 41,
		"Q": 42,
		"P": 43,
		"O": 44,
		"N": 45,
		"M": 46,
		"L": 47,
		"K": 48,
		"J": 49,
		"I": 50,
		"H": 51,
		"G": 52,
		"F": 53,
		"E": 54,
		"D": 55,
		"C": 56,
		"B": 57,
		"A": 58,
	}

	if prefixFlag {
		res = prefixCommands[superFlag]
	}

	for key, value := range convTable {
		s = strings.Replace(s, key, fmt.Sprintf("${@:%d:1}", value), -1)
	}

	for key, value := range numTable {
		s = strings.Replace(s, key, fmt.Sprintf("%s", value), -1)
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
