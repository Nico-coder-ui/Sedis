package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var flushallCmd = &cobra.Command{
	Use:   "FLUSHALL",
	Short: "Use FLUSHALL on Sedis",
	RunE:  flushallHandle,
}

func flushallHandle(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("usage: FLUSHALL")
	}

	req, err := http.NewRequest("POST", "http://sedis:8085/flushall", nil)
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
	rootCmd.AddCommand(flushallCmd)
}
