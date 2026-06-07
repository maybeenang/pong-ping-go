.PHONY: dev-backend dev-frontend dev setup

setup:
	go mod tidy
	cd web && pnpm install

dev-backend:
	air

dev-frontend:
	cd web && pnpm run dev

dev:
	make -j 2 dev-backend dev-frontend
