package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//go:embed template/snippet.go
var snippetGo []byte

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the environment for the atcoder",
	RunE: func(cmd *cobra.Command, args []string) error {
		num := args[0]
		question := args[1]
		lang := args[2]
		if num == "" {
			return fmt.Errorf("num is required")
		}
		if question == "" {
			return fmt.Errorf("question is required")
		}
		template, err := getTemplate(lang)
		if err != nil {
			return fmt.Errorf("failed to get template: %w", err)
		}

		// このgoファイルからの相対パスで指定
		os.Mkdir(fmt.Sprintf("%s/", num), 0755)
		os.Mkdir(fmt.Sprintf("%s/%s/", num, question), 0755)
		file := "%s/%s/main.go"
		os.WriteFile(fmt.Sprintf(file, num, question), template, 0644)
		fmt.Printf("[Template generated] -> %s", file)
		return nil
	},
}

func getTemplate(lang string) ([]byte, error) {
	if lang == "go" {
		return snippetGo, nil
	}
	return nil, fmt.Errorf("language %s is not supported", lang)
}
