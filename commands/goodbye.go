package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GoodbyeCmd = &cobra.Command{
	Use:   "goodbye",
	Short: "Prints goodbye message",
	Long:  `Prints goodbye message to the console`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Goodbye, world!")
	},
}