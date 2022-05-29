package utils

import "crypto/sha256"

func DoubleHash(data []byte) [32]byte {
	hash := sha256.Sum256(data)
	return sha256.Sum256(hash[:])
}
