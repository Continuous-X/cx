package check

import (
	"cx/pkg/output"
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   CommandCheck,
		Short: "check with sub commands",
		Long:  `choose a check`,
		Run: func(cmd *cobra.Command, args []string) {
			message := fmt.Sprintf("use %s", cmd.Name())
			output.PrintCliInfo(message)

		},
	}

	cmd.AddCommand(NewCmdCGhPtotection())
	return cmd
}
