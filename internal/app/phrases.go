package app

import (
	"fmt"
	"regexp"
	"time"
)

const (
	defaultReplyText      = "Робот автоматизации бизнеса приветствует вас"
	defaultDateTimeLayout = "02.01.2006"
)

var createProjectRequest = regexp.MustCompile(`^\s*Создать проект [\w\p{L}-]+ [\w\d]+\s*`)
var listProjiectsRequest = regexp.MustCompile(`^\s*Мои проекты\s*`)

func buildUserCreatedMessage(displayName, login string) string {
	return fmt.Sprintf("Создан новый пользователь: %s. Сгенерированный логин: '%s'.", displayName, login)
}

func buildProcessingFailedMessage(err error) string {
	return fmt.Sprintf("Ошибка при обработке сообщения: %s", err.Error())
}

func buildBuildingReplyFailedMessage(err error) string {
	return fmt.Sprintf("Ошибка при построении ответа: %s", err.Error())
} // FIXME think about func names

func buildCreateProjectSuccessMessage(title string, dueDate time.Time, id int) string {
	return fmt.Sprintf(
		"Создан проект %s с датой завершения %s, ID проекта: %d",
		title, dueDate.Format(time.RubyDate), id,
	)
}

func buildCreateProjectFailedMessage(err error) string {
	return fmt.Sprint("не удалось создать проект: %s", err.Error())
}

func buildListProjectsFailedMessage(err error) string {
	return fmt.Sprintf("не удалось получить список проектов: %s", err.Error())
}

func buildProjectNotificationText(title string, id int, minutes int64) string {
	return fmt.Sprintf("До завершения проекта %s (ID: %d) осталось всего %d минут!", title, id, minutes)
}

func buildNoProjectsText() string {
	return "У вас пока нет проектов"
}

func buildListProjectsHeader() string {
	return "Ваши проекты:\n"
}

func buildListProjectsLine(idx, id int, title string, dueDate time.Time) string {
	return fmt.Sprintf(
		"%d) ID: %d, название: %s, дата завершения: %s\n",
		idx, id, title, dueDate.Format(time.RubyDate),
	)
}
