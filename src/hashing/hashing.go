package hashing

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "math/big"
	"fmt"
)

// Generates a random 32-byte identifier
func CreateId() ([]byte, error) {
    id := make([]byte, 32)
    _, err := rand.Read(id)
    if err != nil {
        return nil, fmt.Errorf("Error generating byte array")
    }
    return id, nil
}

// Computes the SHA-256 hash of a byte array
func GenerateHash(Id []byte) string {
    hasher := sha256.New()
    hasher.Write(Id)
    hash := hasher.Sum(nil)
    hashHex := hex.EncodeToString(hash)
    return hashHex
}

// HashToInt converts a hexadecimal hash string to an *big.Int object.
func HashToInt(hashHex string) (*big.Int, error) {
    intValue, ok := new(big.Int).SetString(hashHex, 16)
    if !ok {
        return nil, fmt.Errorf("invalid hex: %s", hashHex)
    }
    return intValue, nil
}

