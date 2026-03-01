package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stikypiston/satchel/internal"
)

var dropYes bool

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Clear the entire satchel",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !dropYes {
			fmt.Print("Are you sure you want to drop all files from the satchel? [y/N] ")
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(answer))
			if answer != "y" && answer != "yes" {
				fmt.Println("Aborted.")
				return nil
			}
		}

		dir := internal.BaseDir()
		if err := os.RemoveAll(dir); err != nil {
			return err
		}

		fmt.Println("Satchel dropped.")
		return nil
	},
}

func init() {
	dropCmd.Flags().BoolVarP(&dropYes, "yes", "y", false, "skip confirmation prompt")
	rootCmd.AddCommand(dropCmd)
}
