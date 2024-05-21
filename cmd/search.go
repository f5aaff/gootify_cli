package cmd

import "github.com/spf13/cobra"

var searchCmd = cobra.Command{
	Use:   "search",
	Short: "search for items in spotify's catalogue",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		search()
	},
}
