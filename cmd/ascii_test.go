package cmd

import (
	"bufio"
	"os"
	"testing"
)

func pasteUnexpect(s string, t *testing.T) {
}

func TestConvertToASCII(t *testing.T) {
	t.Run("date convert", func(t *testing.T) {

		writeToTMP("date")

		actual := convertToASCII()
		expect := "$'\\x64\\x61\\x74\\x65'\n"

		check(expect, actual, t)

	})
}

func TestConvertToString(t *testing.T) {
	t.Run("$'\\x64\\x61\\x74\\x65' => date", func(t *testing.T) {
		writeToTMP("$'\\x64\\x61\\x74\\x65'")

		actual := convertToString()
		expect := "date\n"

		check(expect, actual, t)

	})
}

func check(expect string, actual string, t *testing.T) {
	if actual != expect {
		t.Errorf("Unexpected result : %s\n\texpect : %s\n", actual, expect)
	}
}

func writeToTMP(s string) {
	file, _ := os.Create(TMP_PATH)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(s)
	writer.Flush()
}
