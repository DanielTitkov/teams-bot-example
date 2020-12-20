package app

import (
	"fmt"
	"regexp"
)

const (
	defaultReplyText = "Робот автоматизации бизнеса приветствует вас"
)

var createProjectRequest = regexp.MustCompile(`Создать проект [\w\d]+ [\w\d]+`)

func buildUserCreatedMessage(displayName, login string) string {
	return fmt.Sprintf("Создан новый пользователь: %s. Сгенерированный логин: '%s'.", displayName, login)
}

func buildProcessingFailedMessage(err error) string {
	return fmt.Sprintf("Ошибка при обработке сообщения %s", err.Error())
}

func buildBuildingReplyFailedMessage(err error) string {
	return fmt.Sprintf("Ошибка при построении ответа %s", err.Error())
} // FIXME think about func names

func buildCreateProjectSuccessMessage(title, dueDate string) string {
	return fmt.Sprintf("Создаю проект %s с датой завершения %s", title, dueDate)
}

func buildCreateProjectFailedMessage() string {
	return "Не удалось создать проект"
}
