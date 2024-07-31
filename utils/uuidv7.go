package utils

import (
	"log"

	"github.com/google/uuid"
)

func NewUUIDV7() string {
	uuidv7, err := uuid.NewV7()
	if err != nil {
		log.Println("error:", err)
		return uuid.NewString()
	}
	return uuidv7.String()
}

func CutStringTo32(toCut string) string {
	return toCut[:32]
}
