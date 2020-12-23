package app

const introCardJSON = (`{
	"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
	"type": "AdaptiveCard",
	"version": "1.0",
	"body": [
	  {
		"type": "TextBlock",
		"text": "Робот автоматизации бизнеса понимает следующие команды:",
		"size": "large",
		"wrap": true
	  },
	  {
		"type": "TextBlock",
		"text": "Создать проект <название> <дата окончания в формате 30.12.2020>",
		"wrap": true,
		"weight": "bolder"
	  },
	  {
		"type": "TextBlock",
		"text": "Пример: Создать проект Мой-Проект_1 30.12.2020",
		"wrap": true
	  },
	  {
		"type": "TextBlock",
		"text": "Мои проекты",
		"wrap": true,
		"weight": "bolder"
	  },
	  {
		"type": "TextBlock",
		"text": "Пример: Мои проекты",
		"wrap": true
	  }	  
	]
  }`)
