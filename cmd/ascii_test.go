package cmd

import (
	"os/exec"
	"testing"
)

func TestConvertTo(t *testing.T) {
	command := exec.Command("echo", "date", ">", TMP_PATH)
	command.Run()

	actual := convertTo()
	expect := "$'\x64\x61\x74\x65'"

}
