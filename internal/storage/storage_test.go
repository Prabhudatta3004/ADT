package storage

import "testing"

func TestKeyValueStore(t *testing.T) {
	store := NewKeyValueStore("test_data.json")

	store.Set("key1", "value1")
	if val, exists := store.Get("key1"); !exists || val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	store.Delete("key1")
	if _, exists := store.Get("key1"); exists {
		t.Errorf("Expected key1 to be deleted")
	}
}
