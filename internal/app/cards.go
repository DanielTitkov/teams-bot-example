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
				"text": "Publish Adaptive Card schema",
				"weight": "bolder",
				"size": "medium"
			  },
			  {
				"type": "ColumnSet",
				"columns": [
				  {
					"type": "Column",
					"width": "auto",
					"items": [
					  {
						"type": "Image",
						"url": "https://pbs.twimg.com/profile_images/3647943215/d7f12830b3c17a5a9e4afcc370e3a37e_400x400.jpeg",
						"size": "small",
						"style": "person"
					  }
					]
				  },
				  {
					"type": "Column",
					"width": "stretch",
					"items": [
					  {
						"type": "TextBlock",
						"text": "Matt Hidinger",
						"weight": "bolder",
						"wrap": true
					  },
					  {
						"type": "TextBlock",
						"spacing": "none",
						"text": "Created {{DATE(2017-02-14T06:08:39Z, SHORT)}}",
						"isSubtle": true,
						"wrap": true
					  }
					]
				  }
				]
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
				  "id": "comment",
				  "isMultiline": true,
				  "placeholder": "Enter your comment"
				}
			  ],
			  "actions": [
				{
				  "type": "Action.Submit",
				  "title": "OK"
				}
			  ]
			}
		  },
		  {
			"type": "Action.OpenUrl",
			"title": "View",
			"url": "https://adaptivecards.io"
		  }
		]
	  }`)
