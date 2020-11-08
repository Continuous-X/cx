package check

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   CommandCheck,
		Short: "check with sub commands",
		Long:  `choose a check`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("use %s", cmd.Name())
		},
	}

	cmd.AddCommand(NewCmdCGhPtotection())
	return cmd
}
