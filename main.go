package main

import (
	"fmt"
	"io"
	"os/exec"
)

func hello() string {
	return "Hello"
}

func xxd(input string) string {
	cmd := exec.Command("xxd")
	stdin, _ := cmd.StdinPipe()
	io.WriteString(stdin, input)
	stdin.Close()
	out, _ := cmd.Output()
	return string(out)
}

func main() {
	fmt.Print(xxd("date"))
}
