package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Read sample.txt in current directory and split by blank lines",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := os.ReadFile("sample.txt")
		if err != nil {
			return fmt.Errorf("failed to read sample.txt: %w", err)
		}

		// Normalize line endings to \n
		normalized := bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
		normalized = bytes.ReplaceAll(normalized, []byte("\r"), []byte("\n"))

		blocks := splitByBlankLines(string(normalized))
		w := bufio.NewWriter(os.Stdout)
		defer w.Flush()

		for i, b := range blocks {
			if strings.TrimSpace(b) == "" {
				continue
			}
			if i > 0 {
				// separator to make block boundaries explicit
				fmt.Fprintln(w, "-----")
			}
			fmt.Fprintln(w, strings.Trim(b, "\n"))
		}
		return nil
	},
}

func splitByBlankLines(s string) []string {
	// Split on one or more blank lines (whitespace-only lines count as blank)
	re := regexp.MustCompile(`(?m)\n[\t\x20]*\n+`)
	parts := re.Split(s, -1)
	return parts
}

func init() {
	rootCmd.AddCommand(runCmd)
}


