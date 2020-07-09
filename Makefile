GOOSE_TARGET_DIR := db/migrations
DATASOURCE := root:@/kanban_development?parseTime=true

run:
	go build -o bin/kanban cmd/kanban/main.go && ./bin/kanban

db/migrate:
	goose -dir "$(GOOSE_TARGET_DIR)" mysql "$(DATASOURCE)" up

db/rollback:
	goose -dir "$(GOOSE_TARGET_DIR)" mysql "$(DATASOURCE)" down

db/status:
	goose -dir "$(GOOSE_TARGET_DIR)" mysql "$(DATASOURCE)" status