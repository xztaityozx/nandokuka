package cmd

import (
	"bufio"
	"os"
	"testing"
)

func TestConvertToASCII(t *testing.T) {
	t.Run("date convert", func(t *testing.T) {

		writeToTMP("date")

		actual := convertToASCII()
		expect := "$'\\x64\\x61\\x74\\x65'\n"

		check(expect, actual, t)

	})

	t.Run("seq 30|awk 'NR%2==0{print}' => ", func(t *testing.T) {
		writeToTMP("seq 30|awk 'NR%2==0{print}'")

		actual := convertToASCII()
		expect := "$'\\x73\\x65\\x71\\x20\\x33\\x30\\x7c\\x61\\x77\\x6b\\x20\\x27\\x4e\\x52\\x25\\x32\\x3d\\x3d\\x30\\x7b\\x70\\x72\\x69\\x6e\\x74\\x7d\\x27'\n"

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
	t.Run("$'\\x73\\x65\\x71\\x20\\x33\\x30\\x7c\\x61\\x77\\x6b\\x20\\x27\\x4e\\x52\\x25\\x32\\x3d\\x3d\\x30\\x7b\\x70\\x72\\x69\\x6e\\x74\\x7d\\x27' =>", func(t *testing.T) {
		writeToTMP("$'\\x73\\x65\\x71\\x20\\x33\\x30\\x7c\\x61\\x77\\x6b\\x20\\x27\\x4e\\x52\\x25\\x32\\x3d\\x3d\\x30\\x7b\\x70\\x72\\x69\\x6e\\x74\\x7d\\x27'")

		actual := convertToString()
		expect := "seq 30|awk 'NR%2==0{print}'\n"

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
