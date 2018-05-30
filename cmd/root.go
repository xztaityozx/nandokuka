// Copyright © 2018 xztaityozx
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nandokuka",
	Short: "Obfuscation tool for ShellGei",
	Long: `Nandokuka tool for ShellGei.
Require:
	- xxd`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

const (
	TMP_PATH string = "/tmp/nandokuka_tmp"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var decodeFlag bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&decodeFlag, "decode", "d", false, "decode data")
}

func redirect(path []string) error {

	var scanner *bufio.Scanner
	switch len(path) {
	case 0:
		scanner = bufio.NewScanner(os.Stdin)
	case 1:
		if path[0] == TMP_PATH {
			return errors.New("could not open " + TMP_PATH)
		}
		file, _ := os.Open(path[0])
		defer file.Close()
		scanner = bufio.NewScanner(file)
	default:
		return errors.New("nandokuka can open only 1 file.")
	}

	tmpFile, _ := os.Create(TMP_PATH)
	defer tmpFile.Close()

	writer := bufio.NewWriter(tmpFile)

	for scanner.Scan() {
		writer.WriteString(scanner.Text())
	}
	writer.Flush()

	return scanner.Err()
}

func surround(pattern bool, s string) string {
	if pattern {
		return "`" + s + "`"
	} else {
		return "$(" + s + ")"
	}
}

func bstring(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func sbyte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func readFromFile(path string, action func(string) string) string {
	var file *os.File
	if len(path) == 0 {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var rt string
	for s.Scan() {
		rt += action(s.Text())
	}
	return rt
}
