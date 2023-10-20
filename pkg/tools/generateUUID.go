package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofrs/uuid"
	"log"
	"os"
	"strings"
)

type UUIDStringGenerator struct {
}

func NewGen() UUIDStringGenerator {
	return UUIDStringGenerator{}
}

func (gen *UUIDStringGenerator) GenerateUUID() string {
	hostname, err := os.Hostname()

	u2, err := uuid.NewV7()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	data := strings.Join([]string{u2.String(), hostname}, "-")

	// Данные манипуляции нужны, чтобы провести data к форме подходящей для UUID
	hash := md5.Sum([]byte(data))
	md5Hex := hex.EncodeToString(hash[:])

	uudiWithHostName, err := uuid.FromString(md5Hex)
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", uudiWithHostName, err)
	}

	return uudiWithHostName.String()
}
