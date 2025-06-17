package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var ttlCmd = &cobra.Command{
	Use:   "TTL",
	Short: "Use TTL on Sedis",
	RunE:  ttlHandle,
}

func ttlHandle(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: TTL key test")
	}

	msg := "TTL " + strings.Join(args, " ")
	req, err := http.NewRequest("POST", "http://sedis:8085/ttl", bytes.NewBufferString(msg))
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
	rootCmd.AddCommand(ttlCmd)
}
