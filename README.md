# AI Agent Manager

## Project Structure

`ai-agent-manager/
cmd/
    agent-manager/  # основной бинарь (gRPC + gateway + воркеры)
    migrate/        # утилита для миграций
api/                # контракты gRPC + события Kafka
    proto/
internal/
    domain/         # бизнес-модель, агрегаты, правила
    use_case/       # сценарии (application layer)
    adapters/       # in/out (grpc, pg, kafka, tools, llm)
    platform/       # pg/kafka init, observability, runtime-воркеры, конфиг
pkg/                # общие утилиты (id, clock, errors, backoff)
migrations/         # SQL-миграции (golang-migrate)
deploy/             # dev-среда
scripts/            # вспомогательные скрипты (генерация, линтеры, dev-tools)
Makefile
README.md
`