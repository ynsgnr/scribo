package common

import (
	"crypto/sha256"
	"fmt"
)

func CalculateIDs(value, internalSecret, externalSecret string) (string, string, error) {
	valueHasher := sha256.New()
	_, err := valueHasher.Write([]byte(value))
	if err != nil {
		return "", "", err
	}
	externalSecretHasher := sha256.New()
	_, err = externalSecretHasher.Write(valueHasher.Sum(nil))
	if err != nil {
		return "", "", err
	}
	_, err = externalSecretHasher.Write([]byte(externalSecret))
	if err != nil {
		return "", "", err
	}
	externalID := fmt.Sprintf("%x", externalSecretHasher.Sum(nil))
	internalSecretHasher := sha256.New()
	_, err = internalSecretHasher.Write([]byte(externalID))
	if err != nil {
		return "", "", err
	}
	_, err = internalSecretHasher.Write([]byte(internalSecret))
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("%x", internalSecretHasher.Sum(nil)), externalID, nil
}

func CalculateInternalID(externalID, internalSecret string) (string, error) {
	internalSecretHasher := sha256.New()
	_, err := internalSecretHasher.Write([]byte(externalID))
	if err != nil {
		return "", err
	}
	_, err = internalSecretHasher.Write([]byte(internalSecret))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", internalSecretHasher.Sum(nil)), nil
}
