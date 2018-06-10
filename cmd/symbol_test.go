package cmd

import (
	"testing"
)

func TestEncodeToSymbolOnly(t *testing.T) {
	t.Run("encode "+TEST_DATE, func(t *testing.T) {
		expect := "${@:2$((1+2)):1}${@:2$((2+2*2)):1}${@:$((2*2*2-1)):1}${@:22:1}"
		actual := encodeToSymbolOnly(TEST_DATE)

		Check(expect, actual, t)
	})
	t.Run("encode prefix "+TEST_DATE, func(t *testing.T) {
		expect := "A=$(. 2>&1);A=${A##*.};${A:1$((2*2*2+1)):1}${A:$((2+2)):1}${A:1$((2*2+2*2)):1} -- {z..A};${@:2$((1+2)):1}${@:2$((2+2*2)):1}${@:$((2*2*2-1)):1}${@:22:1}"
		prefixFlag = true
		actual := encodeToSymbolOnly(TEST_DATE)

		Check(expect, actual, t)
	})

	t.Run("encode prefix "+TEST_SEQ_AWK, func(t *testing.T) {
		prefixFlag = true
		actual := encodeToSymbolOnly(TEST_SEQ_AWK)
		expect := `A=$(. 2>&1);A=${A##*.};${A:1$((2*2*2+1)):1}${A:$((2+2)):1}${A:1$((2*2+2*2)):1} -- {z..A};${@:$((2*2+2*2)):1}${@:22:1}${@:1$((1-1)):1} $((1+2))$((1-1))|${@:2$((2+2*2)):1}${@:$((2+2)):1}${@:1$((2+2*2)):1} ${@:$((2+2))$((2*2+1)):1}${@:$((2+2))1:1}%2==$((1-1)){${@:11:1}${@:$((2*2*2+1)):1}${@:1$((2*2+2*2)):1}${@:1$((1+2)):1}${@:$((2*2*2-1)):1}}`

		Check(expect, actual, t)
	})
	t.Run("encode "+TEST_SEQ_AWK, func(t *testing.T) {
		prefixFlag = false
		actual := encodeToSymbolOnly(TEST_SEQ_AWK)
		expect := `${@:$((2*2+2*2)):1}${@:22:1}${@:1$((1-1)):1} $((1+2))$((1-1))|${@:2$((2+2*2)):1}${@:$((2+2)):1}${@:1$((2+2*2)):1} ${@:$((2+2))$((2*2+1)):1}${@:$((2+2))1:1}%2==$((1-1)){${@:11:1}${@:$((2*2*2+1)):1}${@:1$((2*2+2*2)):1}${@:1$((1+2)):1}${@:$((2*2*2-1)):1}}`

		Check(expect, actual, t)
	})
}
