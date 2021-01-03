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

const testActionCardJSON = (`{
		"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
		"type": "AdaptiveCard",
		"version": "1.0",
		"body": [
		  {
			"type": "Container",
			"items": [
			  {
				"type": "TextBlock",
				"text": "Создать проект",
				"weight": "bolder",
				"size": "medium"
			  }
			]
		  },
		  {
			"type": "Container",
			"items": [
			  {
				"type": "TextBlock",
				"text": "Now that we have defined the main rules and features of the format, we need to produce a schema and publish it to GitHub. The schema will be the starting point of our reference documentation.",
				"wrap": true
			  },
			  {
				"type": "FactSet",
				"facts": [
				  {
					"title": "Board:",
					"value": "Adaptive Card"
				  },
				  {
					"title": "List:",
					"value": "Backlog"
				  },
				  {
					"title": "Assigned to:",
					"value": "Matt Hidinger"
				  },
				  {
					"title": "Due date:",
					"value": "Not set"
				  }
				]
			  }
			]
		  }
		],
		"actions": [
		  {
			"type": "Action.ShowCard",
			"title": "Comment",
			"card": {
			  "type": "AdaptiveCard",
			  "body": [
				{
				  "type": "Input.Text",
				  "id": "title",
				  "placeholder": "Название проекта"
				},
				{
					"type": "Input.Text",
					"id": "dueDate",
					"placeholder": "Дата завершения проекта"
				},
				{
					"type": "Input.Text",
					"id": "app",
					"value": "bar",
					"isVisible": false
				},
				{
					"type": "Input.Text",
					"id": "action",
					"value": "createProject",
					"isVisible": false
				}
			  ],
			  "actions": [
				{
				  "type": "Action.Submit",
				  "title": "Создать проект"
				}
			  ]
			}
		  }
		]
	  }`)
