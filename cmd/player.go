package cmd

import (
	"github.com/spf13/cobra"
)

var playerCmd = cobra.Command{
	Use:   "player",
	Short: "player commands",
	Long: `play: expects a spotify track/album/playlist/episode,etc. ID,
	 if given no args, will resume play.`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		playControls([]*string{&args[0], &args[1]})
	},
}
