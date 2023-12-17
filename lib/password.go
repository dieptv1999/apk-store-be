package lib

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	rn "math/rand"
	"time"
)

type PasswordHash struct {
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (p *PasswordHash) RandStringBytes(n int) string {
	r2 := rn.New(rn.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r2.Intn(len(letterBytes))]
	}
	return string(b)
}

func (p *PasswordHash) GeneratePassword(password string, saltSize int, algorithm string) (string, string, error) {
	// First generate random 16 byte salt
	var salt = generateRandomSalt(saltSize)

	// Hash password using the salt
	return hashPassword(password, salt, algorithm), string(salt), nil
}

func (p *PasswordHash) VerifyPassword(password string, hash string, salt string, algorithm string) bool {
	return doPasswordsMatch(hash, password, []byte(salt), algorithm)
}

// Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a hex string
func hashPassword(password string, salt []byte, algorithm string) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var hasher hash.Hash
	switch algorithm {
	case "SHA256":
		hasher = sha256.New()
	case "SHA1":
		hasher = sha1.New()
	default:
		hasher = sha512.New()
	}

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

// Check if two passwords match
func doPasswordsMatch(hashedPassword, currPassword string,
	salt []byte, algorithm string) bool {
	var currPasswordHash = hashPassword(currPassword, salt, algorithm)

	return hashedPassword == currPasswordHash
}

func generateRandomSalt(saltSize int) []byte {
	r2 := rn.New(rn.NewSource(time.Now().UnixNano()))
	b := make([]byte, saltSize)
	for i := range b {
		b[i] = letterBytes[r2.Intn(len(letterBytes))]
	}
	return b
}

func NewPasswordHash() PasswordHash {
	return PasswordHash{}
}
