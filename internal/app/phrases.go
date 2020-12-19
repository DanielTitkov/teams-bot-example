package app

import "fmt"

func buildUserCreatedMessage(displayName, login string) string {
	return fmt.Sprintf("Создан новый пользователь: %s. Сгенерированный логин: '%s'.", displayName, login)
}
