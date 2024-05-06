package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre Vacio")
	}
	message := fmt.Sprintf(randomSaludos(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messagesMap := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messagesMap[name] = message
	}
	return messagesMap, nil
}

func randomSaludos() string {
	saludos := []string{
		"Hola, %v!, Bienvenido!",
		"Que bueno verte, %v!",
		"Que andas haciendo %v?",
	}
	return saludos[rand.Intn(len(saludos))]
}
