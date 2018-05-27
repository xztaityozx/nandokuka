package cmd

import (
	"testing"
)

func TestConvertToEcho(t *testing.T) {
	t.Run("date => ", func(t *testing.T) {
		actual := convertToEcho(TEST_DATE)
		expect := "$(echo \"d\")$(echo \"a\")`echo \"t\"``echo \"e\"`"

		Check(expect, actual, t)
	})
}
