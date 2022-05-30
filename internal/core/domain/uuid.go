package domain

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.New().String()
}

func IsUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
