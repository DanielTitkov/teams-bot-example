APP := app

.PHONY: run
run: entgen
	go run cmd/$(APP)/main.go

.PHONY: db
db:
	cd deployments/dev && docker-compose up -d --force-recreate --build --remove-orphans --always-recreate-deps --renew-anon-volumes

.PHONY: entgen
entgen:
	cd internal/repository/entgo && go generate ./ent

.PHONY: expose
expose:
	/home/titkovd/Desktop/ngrok http -hostname=teamsbot1.ngrok.io 3978