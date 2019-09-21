package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/sidletsky/sitemap"
	"github.com/sidletsky/sitemap/internal"
)

var file string
var RootCmd = &cobra.Command{
	Use:              "sitemap [URL] [FLAGS]",
	Short:            "sitemap can generate sitemap file for the provided url",
	TraverseChildren: true,
	Args:             cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		sitemap, err := sitemap.Parse(url, nil)
		if err != nil {
			log.Fatal(err)
		}
		internal.CreateFile(file, sitemap)
	},
}

func main() {
	RootCmd.PersistentFlags().StringVarP(&file, "file", "f", "sitemap.xml", "name of an output file")
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
