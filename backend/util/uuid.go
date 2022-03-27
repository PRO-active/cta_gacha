package util

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}

