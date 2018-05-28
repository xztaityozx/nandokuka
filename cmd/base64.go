// Copyright © 2018 xztaityozx
//

package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var JPBASE64_ENCODE_TABLE = map[string]string{
	"a": "ず",
	"b": "せ",
	"c": "ぜ",
	"d": "そ",
	"e": "ぞ",
	"f": "た",
	"g": "だ",
	"h": "ち",
	"i": "ぢ",
	"j": "っ",
	"k": "つ",
	"l": "づ",
	"m": "て",
	"n": "で",
	"o": "と",
	"p": "ど",
	"q": "な",
	"r": "に",
	"s": "ぬ",
	"t": "ね",
	"u": "の",
	"v": "は",
	"w": "ば",
	"x": "ぱ",
	"y": "ひ",
	"z": "び",
	"A": "む",
	"B": "ぁ",
	"C": "あ",
	"D": "ぃ",
	"E": "い",
	"F": "ぅ",
	"G": "う",
	"H": "ぇ",
	"I": "え",
	"J": "ぉ",
	"K": "お",
	"L": "か",
	"M": "が",
	"N": "き",
	"O": "ぎ",
	"P": "く",
	"Q": "ぐ",
	"R": "け",
	"S": "げ",
	"T": "こ",
	"U": "ご",
	"V": "さ",
	"W": "ざ",
	"X": "し",
	"Y": "じ",
	"Z": "す",
}
var JPBASE64_DECODE_TABLE = map[string]string{
	"ず": "a",
	"せ": "b",
	"ぜ": "c",
	"そ": "d",
	"ぞ": "e",
	"た": "f",
	"だ": "g",
	"ち": "h",
	"ぢ": "i",
	"っ": "j",
	"つ": "k",
	"づ": "l",
	"て": "m",
	"で": "n",
	"と": "o",
	"ど": "p",
	"な": "q",
	"に": "r",
	"ぬ": "s",
	"ね": "t",
	"の": "u",
	"は": "v",
	"ば": "w",
	"ぱ": "x",
	"ひ": "y",
	"び": "z",
	"む": "A",
	"ぁ": "B",
	"あ": "C",
	"ぃ": "D",
	"い": "E",
	"ぅ": "F",
	"う": "G",
	"ぇ": "H",
	"え": "I",
	"ぉ": "J",
	"お": "K",
	"か": "L",
	"が": "M",
	"き": "N",
	"ぎ": "O",
	"く": "P",
	"ぐ": "Q",
	"け": "R",
	"げ": "S",
	"こ": "T",
	"ご": "U",
	"さ": "V",
	"ざ": "W",
	"し": "X",
	"じ": "Y",
	"す": "Z",
}

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "base64難読化します",
	Long: `base64難読化コンバーターです
	Usage : nandokuka [-d,--decode] base64 [-j,--jp|-v,--verbose|-e,--excutable] [FILE]

	[FILE]を空にするとstdinから受け取ります

	ex)
	date =>  ZGF0ZQ==

base64エンコードされた文字列は一般に読めないことを利用した難読化です
このサブコマンドは普通に使うとbase64エンコーダーです(base64難読化はそういうもんです)
-e,--excutableをつけると以下のように実行できる形で出力されます

	date => eval $(echo -n ZGF0ZQ==|base64 -d)

-e,--excutableをつけて出力された結果をデコードするときも-e,--excutableをつけてください

-j,--jpを使うと日本語base64難読化します
「ず」をbase64エンコードすると 「44Ga」になりますが、この「a」を利用した難読化になります

	ex)
	date => そずねぞ

この難読化をしたコマンド列をそのまま実行することはできません
日本語base64難読化では-e,--excutableが無効になります
`,
	Run: func(cmd *cobra.Command, args []string) {
		file := ""
		if len(args) == 1 {
			file = args[0]
		}
		res := readFromFile(file, func(s string) string {
			if decodeFlag {
				return decodeToCommandFromBase64(sbyte(s))
			} else {
				return encodeToBase64FromCommand(sbyte(s))
			}
		})

		if !jpFlag && excutableFlag && !decodeFlag {
			res = "eval \"$(echo -n " + res + "|base64 -d)\""
		}

		fmt.Print(res)
	},
}

func encodeToBase64FromCommand(b []byte) string {

	if jpFlag {
		s := bstring(b)
		for key, value := range JPBASE64_ENCODE_TABLE {
			s = strings.Replace(s, key, value, -1)
			if verbFlag {
				fmt.Printf("(%s,%s) => %s\n", key, value, s)
			}
		}
		return s
	} else {
		writeBuf := new(bytes.Buffer)
		writeBuf.Reset()
		encoder := base64.NewEncoder(base64.StdEncoding, writeBuf)
		defer encoder.Close()
		encoder.Write(b)
		encoder.Close()

		res := writeBuf.String()
		if verbFlag {
			fmt.Printf("%s => %s\n", bstring(b), res)
		}

		return res
	}
}

func decodeToCommandFromBase64(b []byte) string {
	if jpFlag {
		r := []rune(bstring(b))
		var rt string
		for i := range r {
			get, ok := JPBASE64_DECODE_TABLE[string(r[i])]
			if ok {
				rt += get
			} else {
				rt += string(r[i])
			}
			if verbFlag {
				fmt.Println(rt)
			}
		}
		return rt
	} else {
		s := bstring(b)
		if excutableFlag {
			s = strings.Replace(s, "eval \"$(echo -n ", "", 1)
			s = strings.Replace(s, "|base64 -d)\"", "", 1)
		}
		dst, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Fatal(err)
		}
		res := bstring(dst)
		if verbFlag {
			fmt.Printf("%s => %s\n", bstring(b), res)
		}
		return res

	}
}

var jpFlag bool
var excutableFlag bool

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().BoolVarP(&verbFlag, "verbose", "v", false, "変換過程も表示します")
	base64Cmd.Flags().BoolVarP(&jpFlag, "jp", "j", false, "日本語base64難読化します")
	base64Cmd.Flags().BoolVarP(&excutableFlag, "excutable", "e", false, "変換結果を実行できる形で出力します。日本語base64難読化では無視されます")
}
