package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/prabhudatta3004/ADT/utils"
)

type KeyValueStore struct {
	data     map[string]interface{}
	mu       sync.RWMutex
	filePath string
}

func NewKeyValueStore(filePath string) *KeyValueStore {
	// If the provided filePath is a directory, append a default filename
	if filePath == "./data" {
		filePath = "./data/adt_data.json"
	}
	store := &KeyValueStore{
		data:     make(map[string]interface{}),
		filePath: filePath,
	}
	store.loadFromFile() // Load data from file on startup
	return store
}

// Load data from file
func (s *KeyValueStore) loadFromFile() {
	// Extract directory path from the filePath
	dir := filepath.Dir(s.filePath)

	// Check if the directory exists, if not, create it
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			utils.Logger.Printf("Failed to create directory: %v", err)
			return
		}
	}

	// Attempt to open the file for loading data
	file, err := os.Open(s.filePath)
	if os.IsNotExist(err) {
		// If the file does not exist, it's okay (e.g., first run), so no error is needed
		utils.Logger.Printf("Data file does not exist, initializing new data.")
		return
	} else if err != nil {
		utils.Logger.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	// Decode data from the file into the store's map
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.data); err != nil {
		utils.Logger.Printf("Failed to decode file: %v", err)
	}
}

// Save data to file
func (s *KeyValueStore) saveToFile() {
	file, err := os.Create(s.filePath)
	if err != nil {
		utils.Logger.Printf("Failed to create file: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(s.data); err != nil {
		utils.Logger.Printf("Failed to encode data to file: %v", err)
	}
}

// Set operation with persistence
func (s *KeyValueStore) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	s.saveToFile()
}

// Get operation
func (s *KeyValueStore) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.data[key]
	return val, exists
}

// Delete operation with persistence
func (s *KeyValueStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	s.saveToFile()
}
