/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	baseURL = "http://localhost:3000/"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gootify_cli",
	Short: "A brief description of your application",
	Long:  `supports basic functions`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(&playerCmd)
	rootCmd.AddCommand(&devicecmd)
	devicecmd.AddCommand(&listcmd)
	rootCmd.AddCommand(&searchCmd)
	searchCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "the query to search with, e.g \"paranoid black sabbath\"")
	searchCmd.PersistentFlags().StringSliceVarP(&types, "types", "T", []string{}, "type of content to search for")
	searchCmd.Flags().IntVarP(&limit, "limit", "l", 0, "maximum number of results to return")
	searchCmd.MarkPersistentFlagRequired("query")

}
