package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version    string
	commitHash string
)

const (
	flagConfigFilePaht = "config-file-path"
)

func main() {
	rootCommand := &cobra.Command{
		Version: fmt.Sprintf("%s (%s)", version, commitHash),
	}
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
	}
}
