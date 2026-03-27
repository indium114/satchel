package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/indium114/satchel/internal"
)

var force bool

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put <id>",
	Short: "Paste a file from the satchel into the current directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return errors.New("Invalid ID")
		}

		idx, err := internal.Load()
		if err != nil {
			return err
		}

		item, ok := idx.Items[id]
		if !ok {
			return errors.New("ID not found")
		}

		src := filepath.Join(internal.ObjectsDir(), fmt.Sprintf("%d", id))
		dest := item.Name

		if !force {
			if _, err := os.Stat(dest); err == nil {
				return errors.New("File already exists (use --force to overwrite)")
			}
		}

		if err := copyFile(src, dest); err != nil {
			return err
		}

		fmt.Printf("Put %s\n", dest)
		return nil
	},
}

func init() {
	putCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing file")
	rootCmd.AddCommand(putCmd)
}
