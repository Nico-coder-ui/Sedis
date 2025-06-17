package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "LIST",
	Short: "Use LIST on Sedis",
	RunE:  listHandle,
}

func listHandle(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("usage: LIST")
	}

	req, err := http.NewRequest("GET", "http://sedis:8085/list", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Server response:", string(body))

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
