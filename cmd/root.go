// Copyright © 2018 xztaityozx
//

package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	TMP_PATH string = "/tmp/nandokuka_tmp"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nandokuka",
	Short: "nandokuka is Nandokuka tool for ShellGei.",
	Long: `nandokuka is Nandokuka tool for ShellGei.
This tool supports several conversion methods. 
if you want more infomation see help.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//cmd.Help()
	//fmt.Print(args)
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
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
