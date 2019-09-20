package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sidletsky/sitemap"
)

var RootCmd = &cobra.Command{
	Use:   "sitemap",
	Short: "sitemap can generate sitemap.xml for the provided url",
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		_, _ = sitemap.Parse(url)
	},
}

func main() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
