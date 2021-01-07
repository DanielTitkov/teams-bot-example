package app

import (
	cards "github.com/DanielTitkov/go-adaptive-cards"
)

func buildIntroCard() *cards.Card {
	c := cards.New([]cards.Node{
		&cards.TextBlock{
			Text:   "Робот автоматизации бизнеса понимает следующие команды:",
			Weight: "bolder",
			Wrap:   cards.TruePtr(),
			Size:   "large",
		},
		&cards.Container{
			Style: "emphasis",
			Items: []cards.Node{
				&cards.TextBlock{
					Text:   "Создать проект <название> <дата окончания в формате 30.12.2020>",
					Weight: "bolder",
					Wrap:   cards.TruePtr(),
				},
				&cards.TextBlock{
					Text: "Пример: Создать проект Мой-Проект_1 30.12.2020",
					Wrap: cards.TruePtr(),
				},
				&cards.TextBlock{
					Text:   "Мои проекты",
					Weight: "bolder",
					Wrap:   cards.TruePtr(),
				},
				&cards.TextBlock{
					Text: "Пример: Мои проекты",
					Wrap: cards.TruePtr(),
				},
			},
		},
		&cards.TextBlock{
			Text: "Введите команды текстом или воспользуйтесь конструктором, нажав нужную кнопку.",
			Wrap: cards.TruePtr(),
			Size: "small",
		},
	}, []cards.Node{
		&cards.ActionSubmit{
			Title: "Создать проект",
			Data: map[string]interface{}{
				"app":    "bar",
				"action": "initCreateProject",
			},
		},
		&cards.ActionSubmit{
			Title: "Мои проекты",
			Data: map[string]interface{}{
				"app":    "bar",
				"action": showProjectsAction,
			},
		},
	}).
		WithSchema(cards.DefaultSchema).
		WithVersion(cards.Version12)

	return c
}

func buildCreateProjectCard() *cards.Card {
	c := cards.New([]cards.Node{
		&cards.Container{
			Items: []cards.Node{
				&cards.TextBlock{
					Text: "Название проекта",
				},
				&cards.InputText{
					ID:          "title",
					Placeholder: "Мой проект",
					IsRequired:  cards.TruePtr(),
				},
				&cards.TextBlock{
					Text: "Дата завершения",
				},
				&cards.InputText{
					ID:          "dueDate",
					Placeholder: "30.10.2021",
					IsRequired:  cards.TruePtr(),
				},
			},
		},
	}, []cards.Node{
		&cards.ActionSubmit{
			Title: "Создать!",
			Data: map[string]interface{}{
				"app":    "bar",
				"action": createProjectAction,
			},
		},
	}).
		WithSchema(cards.DefaultSchema).
		WithVersion(cards.Version12)

	return c
}
