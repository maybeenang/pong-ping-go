.PHONY: dev-backend dev-frontend dev setup migrate-up migrate-down migrate-drop migrate-version

setup:
	go mod tidy
	cd web && pnpm install

dev-backend:
	air

dev-frontend:
	cd web && pnpm run dev

dev:
	make -j 2 dev-backend dev-frontend

migrate-up:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down

migrate-drop:
	go run ./cmd/migrate drop

migrate-version:
	go run ./cmd/migrate version
