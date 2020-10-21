help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev-up: ## 開発環境の起動
	docker-compose up -d

dev-db-up: ## 開発用DBを起こす
	docker-compose up -d db

dev-app-watch: ## アプリログの監視
	docker-compose logs -f app

dev-db-watch: ## dbのログを監視
	docker-compose logs -f db

dev-down: ## 開発環境を削除
	docker-compose down --rmi local --volumes
