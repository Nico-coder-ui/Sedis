package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "LOAD",
	Short: "Use LOAD on Sedis",
	RunE:  loadHandle,
}

func loadHandle(cmd *cobra.Command, args []string) error {
	if len(args) != 0 && len(args) != 1 {
		return fmt.Errorf("usage: LOAD [filename]")
	}

	msg := "LOAD " + strings.Join(args, " ")
	req, err := http.NewRequest("POST", "http://sedis:8085/load", bytes.NewBufferString(msg))
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
	rootCmd.AddCommand(loadCmd)
}
