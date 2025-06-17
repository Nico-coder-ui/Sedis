package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "DEL",
	Short: "Use DEL on Sedis",
	RunE:  delHandle,
}

func delHandle(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: DEL key")
	}

	msg := "DEL " + strings.Join(args, " ")
	req, err := http.NewRequest("POST", "http://sedis:8085/del", bytes.NewBufferString(msg))
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
	rootCmd.AddCommand(delCmd)
}
