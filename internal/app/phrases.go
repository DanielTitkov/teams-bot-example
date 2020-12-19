package app

import "fmt"

const (
	defaultReplyText = "Робот автоматизации бизнеса приветствует вас"
)

func buildUserCreatedMessage(displayName, login string) string {
	return fmt.Sprintf("Создан новый пользователь: %s. Сгенерированный логин: '%s'.", displayName, login)
}

func buildProcessingFailedMessage(err error) string {
	return fmt.Sprintf("Ошибка при обработке сообщения %s", err.Error())
}

func buildBuildingReplyFailedMessage(err error) string {
	return fmt.Sprintf("Ошибка при построении ответа %s", err.Error())
}
