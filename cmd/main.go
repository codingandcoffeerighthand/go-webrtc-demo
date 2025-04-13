package main

import (
	"fmt"

	"github.com/codingandcoffeerighthand/go-webrtc-demo/internal/app"
	"github.com/spf13/cobra"
)

var (
	version    string
	commitHash string
)

const (
	flagConfigFilePaht = "config-file-path"
)

var appCommand = &cobra.Command{
	Use:   "run",
	Short: "A demo for go-webrtc",
	Long:  `A demo for go-webrtc`,
	Run: func(cmd *cobra.Command, args []string) {
		var app = app.NewApp("localhost", 8000, "secretKey")
		app.Start()
	},
}

func main() {
	rootCommand := &cobra.Command{
		Version: fmt.Sprintf("%s (%s)", version, commitHash),
	}
	rootCommand.AddCommand(appCommand)
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
	}
}
