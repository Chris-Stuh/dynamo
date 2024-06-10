package tests

import (
    "crypto/sha256"
    "dynamo/src/hashing"
    "encoding/hex"
    "testing"
)

// Verify 32 byte ID was created
func TestCreateId(t *testing.T) {
    newId, err := hashing.CreateId()
    if err != nil {
        t.Fatalf("Error creating ID: %v", err) // Stop if the ID can't be created
    }
    if len(newId) != 32 {
        t.Errorf("Invalid ID: Found length: %d, expected 32", len(newId))
    }
}

// Verify hash works as expected
func TestHash(t *testing.T) {
    newId, err := hashing.CreateId()
    if err != nil {
        t.Fatalf("Error creating ID: %v", err) // Stop if the ID can't be created
    }
    foundHash := hashing.GenerateHash(newId)
    hasher := sha256.New()
    hasher.Write(newId)
    hash := hasher.Sum(nil)
    expectedHash := hex.EncodeToString(hash)
    if foundHash != expectedHash {
        t.Errorf("Hash mismatch: got %s, want %s", foundHash, expectedHash)
    }
}

// Verify an int is returned
func TestIntConversion(t *testing.T) {
    newId, err := hashing.CreateId()
    if err != nil {
        t.Fatalf("Error creating ID: %v", err) // Stop if the ID can't be created
    }
    foundHash := hashing.GenerateHash(newId)
    intVal, err := hashing.HashToInt(foundHash)
    if err != nil {
        t.Errorf("Error converting hash to int: %v", err)
    }
    if intVal == nil {
        t.Errorf("Converted int is nil")
    }
}
