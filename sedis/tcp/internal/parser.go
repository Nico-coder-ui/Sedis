package internal

import (
	"net"
	"sedis/stokage"
	"strings"
)

func ParseMsg(msg string, s *stokage.Store, conn net.Conn) {
	tokens := strings.Fields(msg)

	if len(tokens) == 0 {
		conn.Write([]byte("Empty command\n"))
		return
	}

	cmd := strings.ToUpper(tokens[0])

	switch cmd {
	case "PING":
		conn.Write([]byte("PONG\n"))

	case "HELP":
		HelpHandle(conn)

	case "TTL":
		TtlHandle(tokens, conn, s)

	case "EXISTS":
		ExistsHandle(tokens, conn, s)

	case "LIST":
		ListHandle(tokens, conn, s)

	case "SET":
		SetHandle(tokens, conn, s)

	case "GET":
		GetHandle(tokens, conn, s)

	case "DEL":
		DelHandle(tokens, conn, s)

	case "FLUSHALL":
		FlushallHandle(tokens, conn, s)

	case "SAVE":
		SaveHandle(tokens, conn, s)

	case "LOAD":
		LoadHandle(tokens, conn, s)

	default:
		conn.Write([]byte("Unknown command\n"))
	}
}
