package check

import (
	"cx-installer/pkg/github"
	"cx-installer/pkg/messages"
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdCGhPtotection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   CommandGhProtection,
		Short: "check github protection",
		Long: `compliance check for github repository protection`,
		Run: func(cmd *cobra.Command, args []string) {
			messages.PrintLogfileAndConsole(fmt.Sprintf("use %s", cmd.Name()))

			compliantState := github.GHRepository{
				Organisation:   "",
				RepositoryName: "",
				GhToken:        "2e356175866ff607a7b460d0f46859b4272e261f",
			}.IsCompliant()

			messages.PrintLogfileAndConsole(fmt.Sprintf("compliant state: %t", compliantState))
		},
	}
	return cmd
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

