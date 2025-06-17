package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var existsCmd = &cobra.Command{
	Use:   "EXISTS",
	Short: "Use EXISTS on Sedis",
	RunE:  existsHandle,
}

func existsHandle(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: EXISTS key")
	}

	msg := "EXISTS " + strings.Join(args, " ")
	req, err := http.NewRequest("GET", "http://sedis:8085/exists", bytes.NewBufferString(msg))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Server response: %s\n", string(respBody))

	return nil
}

func init() {
	rootCmd.AddCommand(existsCmd)
}
