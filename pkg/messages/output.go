package messages

import (
	"fmt"
	"log"
)

func PrintLogfileAndConsole(message string)  {
	PrintConsole(message)
	PrintLogfile(message)
}

func PrintLogfile(message string)  {
	log.Print(message)
}

func PrintConsole(message string)  {
	fmt.Print(message)
}
