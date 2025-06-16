package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "SET",
	Short: "Set a data on Sedis",
	RunE:  setHandle,
}

func setHandle(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: SET key value [NX|XX] [EX seconds]")
	}

	msg := "SET " + strings.Join(args, " ")
	fmt.Println("On sedis_manager:")
	fmt.Println(msg)
	resp, err := http.Post("http://sedis:8085/set", "text/plain", bytes.NewBufferString(msg))

	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Server response: %s\n", string(body))

	return nil
}

func init() {
	rootCmd.AddCommand(setCmd)
}
