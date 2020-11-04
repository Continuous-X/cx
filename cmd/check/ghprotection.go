package check

import (
	"cx-installer/pkg/metrics"
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdCGhPtotection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   CommandGhProtection,
		Short: "check github protection",
		Long: `compliance check for github repository protection`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("use %s", cmd.Name())

			metrics.MetricGHRepositoryProtection{
				CliCommand:          cmd.Name(),
				GhProtectionActive:  false,
				GhPullrequestActive: false,
				GhStatusCheckActive: false,
			}.WriteMetric()
		},
	}

	return cmd
}
