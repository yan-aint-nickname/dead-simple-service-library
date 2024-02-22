# Dead Simple Service Library

TODO:

- components:
  - redis (opt-out)
  - postgres (opt-out)
    - driver ??
    - query builder ?? buildsqlx/sqlx (opt-out)
    - migration managment ?? (opt-out)
  - kafka (opt-out)
  - mongodb (opt-out)
  - http(fiber?) (opt-out/obligatory?) https://github.com/gofiber/fiber ~> fastest? + feature-rich
    - use viper as config manager
  - worker/job scheduler (opt-in/opt-out?) https://github.com/go-co-op/gocron?tab=readme-ov-file
  - graphql ?? (opt-in) https://github.com/graph-gophers/graphql-go

- telemetry:
  - opentelemetry (obligatory) https://opentelemetry.io/docs/languages/go/getting-started/
    - https://docs.gofiber.io/contrib/fibersentry_v1.x.x/otelfiber/
  - sentry-sdk (obligatory) https://github.com/getsentry/sentry-go ~> https://docs.gofiber.io/contrib/fibersentry/

- json https://github.com/goccy/go-json ~> fastest + encoding/json compatible

- logging https://github.com/sirupsen/logrus ~> loguru-like + std compatible

- static analysis:
  - golangci-lint (obligatory/opt-out)
  - gocritic (obligatory/opt-out)
  - gosec

- test
  - go test -v -timeout 30s -coverprofile=cover.out -cover (obligatory/opt-out)
  - go tool cover -func=cover.out (obligatory/opt-out)

- pre-commit:
  - pre-commit (global/opt-out)

- local development:
  - replace github.com/user/library => /local/path/library (replace directive works)
