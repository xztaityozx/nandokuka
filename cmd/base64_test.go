package cmd

import (
	"testing"
)

func TestEncode(t *testing.T) {
	t.Run("base64 "+TEST_DATE, func(t *testing.T) {
		actual := encodeToBase64FromCommand(sbyte(TEST_DATE))
		expect := "ZGF0ZQ=="

		Check(expect, actual, t)
	})
	t.Run("base64 "+TEST_SEQ_AWK, func(t *testing.T) {
		actual := encodeToBase64FromCommand(sbyte(TEST_SEQ_AWK))
		expect := "c2VxIDMwfGF3ayAnTlIlMj09MHtwcmludH0n"

		Check(expect, actual, t)
	})
	t.Run("JPbase64 "+TEST_DATE, func(t *testing.T) {
		jpFlag = true
		actual := encodeToBase64FromCommand(sbyte(TEST_DATE))
		expect := "そずねぞ"

		Check(expect, actual, t)
	})
	t.Run("JPbase64 "+TEST_SEQ_AWK, func(t *testing.T) {
		jpFlag = true
		actual := encodeToBase64FromCommand(sbyte(TEST_SEQ_AWK))
		expect := "ぬぞな 30|ずばつ 'きけ%2==0{どにぢでね}'"

		Check(expect, actual, t)
	})

}

func TestDecode(t *testing.T) {
	t.Run("JPbase64 "+TEST_DATE, func(t *testing.T) {
		jpFlag = true
		actual := decodeToCommandFromBase64(sbyte("そずねぞ"))
		expect := TEST_DATE

		Check(expect, actual, t)
	})
	t.Run("JPbase64 "+TEST_SEQ_AWK, func(t *testing.T) {
		jpFlag = true
		actual := decodeToCommandFromBase64(sbyte("ぬぞな 30|ずばつ 'きけ%2==0{どにぢでね}'"))
		expect := TEST_SEQ_AWK

		Check(expect, actual, t)
	})
	t.Run("base64 "+TEST_DATE, func(t *testing.T) {
		jpFlag = false
		actual := decodeToCommandFromBase64(sbyte("ZGF0ZQ=="))
		expect := TEST_DATE

		Check(expect, actual, t)
	})
	t.Run("base64 "+TEST_SEQ_AWK, func(t *testing.T) {
		jpFlag = false
		actual := decodeToCommandFromBase64(sbyte("c2VxIDMwfGF3ayAnTlIlMj09MHtwcmludH0n"))
		expect := TEST_SEQ_AWK

		Check(expect, actual, t)
	})
}
