package internal

import (
	"net"
	"sedis/stokage"
	"strconv"
	"strings"
	"time"
)

func HelpHandle(conn net.Conn) {
	help := "Usage:\n"
	help += "   TTL    key\t\ttest\n"
	help += "   EXISTS key\n"
	help += "   LIST\n"
	help += "   SET key value [NX|XX] [EX seconds]\n"
	help += "   GET    key\n"
	help += "   DEL    key\n"
	help += "   FLUSHALL\n"
	help += "   SAVE   [fileName]\n"
	help += "   LOAD   [fileName]\n"

	conn.Write([]byte(help))
}

func TtlHandle(tokens []string, conn net.Conn, s *stokage.Store) {
	if len(tokens) < 2 {
		conn.Write([]byte("Usage: TTL key\n"))
		return
	}

	key := tokens[1]

	value := s.Ttl(key)

	conn.Write([]byte(value + "\n"))
}

func ExistsHandle(tokens []string, conn net.Conn, s *stokage.Store) {
	if len(tokens) < 2 {
		conn.Write([]byte("Usage: EXISTS key\n"))
		return
	}

	key := tokens[1]

	value := s.Exists(key)

	conn.Write([]byte(value + "\n"))
}

func ListHandle(tokens []string, conn net.Conn, s *stokage.Store) {
	result := s.List()
	if result != "" {
		conn.Write([]byte(result))
	} else {
		conn.Write([]byte("Failed access to data\n"))
	}
}

func SetHandle(tokens []string, conn net.Conn, s *stokage.Store) {
	if len(tokens) < 3 {
		conn.Write([]byte("Usage: SET key value [NX|XX] [EX seconds]\n"))
		return
	}

	key := tokens[1]
	value := tokens[2]

	var expireInSec string
	var hasExpire bool

	for i := 3; i < len(tokens); i++ {
		switch strings.ToUpper(tokens[i]) {
		case "EX":
			if i+1 < len(tokens) {
				expireInSec = tokens[i+1]
				hasExpire = true
				i++
			} else {
				conn.Write([]byte("Missing value for EX\n"))
				return
			}

		case "NX":
			if s.Exists(key) != "0" {
				conn.Write([]byte("\n"))
				return
			}

		case "XX":
			if s.Exists(key) == "0" {
				conn.Write([]byte("\n"))
				return
			}

		default:
			conn.Write([]byte("Unknown modifier: " + tokens[i] + "\n"))
			return
		}
	}

	if hasExpire {
		seconds, err := strconv.Atoi(expireInSec)
		if err != nil {
			conn.Write([]byte("Invalid EX value\n"))
			return
		}
		expireAt := time.Now().Add(time.Duration(seconds) * time.Second)
		s.Expire(key, expireAt)
	}

	s.Set(key, value)
	conn.Write([]byte("OK\n"))
}

func GetHandle(tokens []string, conn net.Conn, s *stokage.Store) {
	if len(tokens) < 2 {
		conn.Write([]byte("Usage: GET key\n"))
		return
	}
	key := tokens[1]

	value, ok := s.Get(key)
	if ok {
		conn.Write([]byte(value + "\n"))
	} else {
		conn.Write([]byte("Key not found\n"))
	}
}

func DelHandle(tokens []string, conn net.Conn, s *stokage.Store) {
	if len(tokens) < 2 {
		conn.Write([]byte("Usage: DEL key\n"))
		return
	}
	key := tokens[1]
	b := s.Del(key)

	if b {
		conn.Write([]byte("OK\n"))
	} else {
		conn.Write([]byte("Key not found\n"))
	}
}

func FlushallHandle(letokens []string, conn net.Conn, s *stokage.Store) {
	ok := s.Flushall()

	if ok {
		conn.Write([]byte("OK\n"))
	} else {
		conn.Write([]byte("Failed to clear all data\n"))
	}
}

func SaveHandle(tokens []string, conn net.Conn, s *stokage.Store) {
	fileName := "stokage.store.json"

	if len(tokens) >= 2 {
		fileName = tokens[1]
		if !strings.HasSuffix(fileName, ".json") {
			conn.Write([]byte("File name must be a .json\n"))
			return
		}
	}

	ok := s.Save(fileName)
	if ok {
		conn.Write([]byte("OK\n"))
	} else {
		conn.Write([]byte("Failed to save data\n"))
	}
}

func LoadHandle(tokens []string, conn net.Conn, s *stokage.Store) {
	fileName := "store.json"

	if len(tokens) >= 2 {
		fileName = tokens[1]
		if !strings.HasSuffix(fileName, ".json") {
			conn.Write([]byte("File name must be a .json\n"))
			return
		}
	}
	ok := s.Load(fileName)
	if ok {
		conn.Write([]byte("OK\n"))
	} else {
		conn.Write([]byte("Failed to load file\n"))
	}
}
