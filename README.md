# AI Agent Manager

AI Agent Manager — это сервис для управления агентами искусственного интеллекта, обеспечивающий взаимодействие через gRPC и обработку событий с помощью Kafka. Проект построен на языке Go и использует PostgreSQL для хранения данных, а также интегрируется с внешними сервисами и инструментами для миграций и оркестрации.

## Project Structure

```plaintext
ai-agent-manager/
├── cmd/
│   ├── agent-manager/   # основной бинарь (gRPC + gateway + воркеры)
│   └── migrate/         # утилита для миграций
├── api/                 # контракты gRPC + события Kafka
│   └── proto/
├── internal/
│   ├── domain/          # бизнес-модель, агрегаты, правила
│   ├── use_case/        # сценарии (application layer)
│   ├── adapters/        # in/out (grpc, pg, kafka, tools, llm)
│   └── platform/        # pg/kafka init, observability, runtime-воркеры, конфиг
├── pkg/                 # общие утилиты (id, clock, errors, backoff)
├── migrations/          # SQL-миграции (golang-migrate)
├── deploy/              # dev-среда
├── scripts/             # вспомогательные скрипты (генерация, линтеры, dev-tools)
├── Makefile
└── README.md
```

## Requirements

- Go 1.21+
- Docker + Docker Compose
- golang-migrate
- Make

## Installation

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/yourusername/ai-agent-manager.git
   cd ai-agent-manager
   ```
2. Установите зависимости:
   ```bash
   make deps
   ```
3. Запустите проект:
   ```bash
   make run
   ```

## Testing

Для запуска тестов выполните команду:
```bash
make test
```