package error

import "cx/pkg/output"

func FailHandleCommand(err error)  {
	if err != nil {
		output.PrintCliError(err)
	}
}