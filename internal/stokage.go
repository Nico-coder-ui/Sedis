package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Store struct {
	Data     map[string]string
	Expiries map[string]time.Time
	Mutex    sync.RWMutex
}

type PersistentStore struct {
	Data     map[string]string    `json:"data"`
	Expiries map[string]time.Time `json:"expiries"`
}

//---------------------------------------------------------------

func (s *Store) Ttl(key string) string {
	s.Mutex.RLock()
	_, ok := s.Data[key]
	s.Mutex.RUnlock()

	if ok {
		remaining := time.Until(s.Expiries[key]).Seconds()
		if remaining <= 0 {
			return strconv.Itoa(-1)
		}
		return fmt.Sprintf("%.0f", remaining)
	} else {
		return strconv.Itoa(-2)
	}
}

func (s *Store) Exists(key string) string {
	s.Mutex.RLock()
	_, ok := s.Data[key]
	s.Mutex.RUnlock()

	if ok {
		return "1"
	}
	return "0"
}

func (s *Store) List() string {
	var result string

	for k, v := range s.Data {
		result += fmt.Sprintf("Key: %s Data: %s", k, v)
		status := s.Ttl(k)
		if status != "-1" && status != "-2" {
			result += fmt.Sprintf(" Expire in: %ss", status)
		}
		result += fmt.Sprintln("")
	}
	return result
}

func (s *Store) Set(key, value string) {
	s.Mutex.Lock()
	s.Data[key] = value
	s.Mutex.Unlock()
}

func (s *Store) Expire(key string, expiration time.Time) {
	s.Mutex.Lock()
	s.Expiries[key] = expiration
	s.Mutex.Unlock()
}

func (s *Store) Get(key string) (string, bool) {
	s.Mutex.RLock()
	value, ok := s.Data[key]
	s.Mutex.RUnlock()

	return value, ok
}

func (s *Store) Del(key string) bool {
	s.Mutex.Lock()
	if _, exists := s.Data[key]; exists {
		delete(s.Data, key)
		s.Mutex.Unlock()
		return true
	}

	s.Mutex.Unlock()
	return false
}

func (s *Store) Save(fileName string) bool {
	s.Mutex.RLock()
	ps := PersistentStore{
		Data:     make(map[string]string, len(s.Data)),
		Expiries: make(map[string]time.Time, len(s.Expiries)),
	}

	for k, v := range s.Data {
		ps.Data[k] = v
	}
	for k, v := range s.Expiries {
		ps.Expiries[k] = v
	}
	s.Mutex.RUnlock()

	data, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		return false
	}

	err = os.WriteFile(fileName, data, 0644)
	return err == nil
}

func (s *Store) Load(fileName string) bool {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return false
	}

	var ps PersistentStore
	err = json.Unmarshal(data, &ps)
	if err != nil {
		return false
	}

	s.Mutex.Lock()
	s.Data = make(map[string]string, len(ps.Data))
	s.Expiries = make(map[string]time.Time, len(ps.Expiries))

	for k, v := range ps.Data {
		s.Data[k] = v
	}
	for k, v := range ps.Expiries {
		s.Expiries[k] = v
	}
	s.Mutex.Unlock()

	return true
}

func (s *Store) Flushall() bool {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Data = make(map[string]string)
	s.Expiries = make(map[string]time.Time)

	if len(s.Data) != 0 || len(s.Expiries) != 0 {
		return false
	}

	return true
}

func (s *Store) StartExpirationLoop() {
	go func() {
		for {
			time.Sleep(time.Second)
			now := time.Now()

			s.Mutex.Lock()
			for key, expiry := range s.Expiries {
				if now.After(expiry) {
					delete(s.Data, key)
					delete(s.Expiries, key)
					fmt.Printf("La donnée '%s' a expirée\n", key)
				}
			}
			s.Mutex.Unlock()
		}
	}()
}

func NewStore() Store {
	return Store{
		Data:     make(map[string]string),
		Expiries: make(map[string]time.Time),
	}
}
