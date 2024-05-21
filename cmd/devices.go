package cmd

import "github.com/spf13/cobra"

var ()

var devicecmd = cobra.Command{
	Use:   "device",
	Short: "device commands",
	Long: `device: used to manage and view devices:
		list: show currently available devices
		`,
	//		change: change to given device, expects device ID
	//		default: if no arguments are given, the current device info will be shown.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			deviceList()
		}
	},
}

var listcmd = cobra.Command{
	Use:   "list",
	Short: "listDevices",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		deviceList()
	},
}
