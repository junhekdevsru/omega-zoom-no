.PHONY: proto clean-proto help

# Генерация protobuf файлов
proto: clean-proto
	@echo "Генерация Go кода из protobuf файлов..."
	@mkdir -p pkg/generate/contracts
	find api/proto -name "*.proto" | xargs protoc --proto_path=api/proto \
		--go_out=pkg/generate/contracts \
		--go_opt=paths=source_relative
	@echo "Генерация protobuf файлов завершена!"

# Очистка сгенерированных файлов
clean-proto:
	@echo "Очистка старых сгенерированных файлов..."
	rm -rf pkg/generate/contracts

# Помощь
help:
	@echo "Доступные команды:"
	@echo "  make proto       - Генерация Go кода из protobuf файлов"
	@echo "  make clean-proto - Удаление сгенерированных protobuf файлов"
	@echo "  make help        - Показать эту справку"
